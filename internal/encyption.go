package internal

import (
	"archive/tar"
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func EncryptAll(config Config) error {
	enigma, err := NewEnigma(config.Key)
	if err != nil {
		return fmt.Errorf("encrypting all paths: %w", err)
	}

	for _, p := range config.Paths {
		stat, err := os.Stat(p)
		if err != nil {
			log.Printf("ERROR: while getting information on file %s : %v", p, err)
		}
		if stat.IsDir() {
			err := enigma.EncryptDirectory(p)
			if err != nil {
				log.Printf("ERROR: while encrypting directory %s : %v", p, err)
			}
		} else {
			err := enigma.EncryptFile(p)
			if err != nil {
				log.Printf("ERROR: while encrypting file %s : %v", p, err)
			}
		}
	}
	return nil
}

func (e *Enigma) EncryptFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read file err: %v", err)
	}
	encrypted, err := e.Encrypt(content)
	if err != nil {
		return fmt.Errorf("creating encrypted file from %s : %w", filePath, err)
	}

	newFilePath := fmt.Sprintf("%s.roll", filePath)
	f, err := os.Create(newFilePath)
	if err != nil {
		return fmt.Errorf("creating encrypted file %s : %w", newFilePath, err)
	}
	defer f.Close()

	_, err = f.Write(encrypted)
	if err != nil {
		return fmt.Errorf("writing to encrypted file %s : %w", newFilePath, err)
	}
	return nil
}

func (e *Enigma) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, e.gcm.NonceSize())
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, fmt.Errorf("initializing nonce: %w", err)
	}
	return e.gcm.Seal(nonce, nonce, data, nil), nil
}

func (e *Enigma) EncryptDirectory(directoryPath string) error {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	err := filepath.Walk(directoryPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			fmt.Println(path, info.Size())

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			hdr := &tar.Header{
				Name: path,
				Mode: 0600,
				Size: info.Size(),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				log.Fatal(err)
			}
			if _, err := tw.Write(content); err != nil {
				log.Fatal(err)
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("creating encrypted file from directory %s : %w", directoryPath, err)
	}

	encryptedTar, err := e.Encrypt(buf.Bytes())
	if err != nil {
		return fmt.Errorf("creating encrypted file from directory %s : %w", directoryPath, err)
	}

	newFilePath := fmt.Sprintf("%s.roll", directoryPath)
	f, err := os.Create(newFilePath)
	if err != nil {
		return fmt.Errorf("creating encrypted file %s : %w", newFilePath, err)
	}
	defer f.Close()

	_, err = f.Write(encryptedTar)
	if err != nil {
		return fmt.Errorf("writing to encrypted file %s : %w", newFilePath, err)
	}
	return nil
}
