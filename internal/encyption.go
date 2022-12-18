package internal

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
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
		if !stat.IsDir() {
			enryptedFilePath := fmt.Sprintf("%s.roll", p)
			err := enigma.EncryptFile(p, enryptedFilePath)
			if err != nil {
				log.Printf("ERROR: while encrypting file %s : %v", p, err)
			}
		}
	}
	return nil
}

func (e *Enigma) EncryptFile(filePath string, encryptedFilePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read file err: %v", err)
	}
	encrypted, err := e.Encrypt(content)
	if err != nil {
		return fmt.Errorf("creating encrypted file from %s : %w", filePath, err)
	}

	f, err := os.Create(encryptedFilePath)
	if err != nil {
		return fmt.Errorf("creating encrypted file %s : %w", encryptedFilePath, err)
	}
	defer f.Close()

	_, err = f.Write(encrypted)
	if err != nil {
		return fmt.Errorf("writing to encrypted file %s : %w", encryptedFilePath, err)
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
