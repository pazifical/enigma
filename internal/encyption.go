package internal

import (
	"archive/tar"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Encrypter struct {
	Gcm   cipher.AEAD
	Nonce []byte
}

func NewEncrypter(key string) (Encrypter, error) {
	key, err := validateEncryptionKey(key)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return Encrypter{}, fmt.Errorf("initializing cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return Encrypter{}, fmt.Errorf("initializing cipher GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return Encrypter{}, fmt.Errorf("initializing nonce: %w", err)
	}
	return Encrypter{
		Gcm:   gcm,
		Nonce: nonce,
	}, nil
}

func EncryptAll(config Config) error {
	encrypter, err := NewEncrypter(config.Key)
	if err != nil {
		return fmt.Errorf("encrypting all paths: %w", err)
	}

	for _, p := range config.Paths {
		stat, err := os.Stat(p)
		if err != nil {
			log.Printf("ERROR: while getting information on file %s : %v", p, err)
		}
		if stat.IsDir() {
			err := encrypter.EncryptDirectory(p)
			if err != nil {
				log.Printf("ERROR: while encrypting directory %s : %v", p, err)
			}
		} else {
			err := encrypter.EncryptFile(p)
			if err != nil {
				log.Printf("ERROR: while encrypting file %s : %v", p, err)
			}
		}
	}
	return nil
}

func validateEncryptionKey(key string) (string, error) {
	if len(key) > 32 {
		return key, fmt.Errorf("maximum AES key length exceeded. %d > 32", len(key))
	}
	for {
		if len(key) == 16 || len(key) == 24 || len(key) == 32 {
			return key, nil
		}
		key = key + "0"
	}
}

func (enc *Encrypter) EncryptFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read file err: %v", err)
	}
	encryptedContent := enc.Gcm.Seal(enc.Nonce, enc.Nonce, content, nil)

	newFilePath := fmt.Sprintf("%s.roll", filePath)
	f, err := os.Create(newFilePath)
	if err != nil {
		return fmt.Errorf("creating encrypted file %s : %w", newFilePath, err)
	}
	defer f.Close()

	_, err = f.Write(encryptedContent)
	if err != nil {
		return fmt.Errorf("writing to encrypted file %s : %w", newFilePath, err)
	}
	return nil
}

func (enc *Encrypter) EncryptDirectory(directoryPath string) error {
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
		log.Fatal(err)
	}

	encryptedTar := enc.Gcm.Seal(enc.Nonce, enc.Nonce, buf.Bytes(), nil)

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
