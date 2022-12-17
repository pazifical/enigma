package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func EncryptAll(config Config) {
	for _, p := range config.Paths {
		stat, err := os.Stat(p)
		if err != nil {
			log.Printf("ERROR: while getting information on file %s : %v", p, err)
		}
		if stat.IsDir() {
			// TODO: implement
		} else {
			err := encryptFile(p, config.Key)
			if err != nil {
				log.Printf("ERROR: while encrypting file %s : %v", p, err)
			}
		}
	}
}

func encryptFile(fPath string, key string) error {
	plainText, err := os.ReadFile(fPath)
	if err != nil {
		log.Fatalf("read file err: %v", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("cipher err: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	log.Println(cipherText)

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return fmt.Errorf("nonce  err: %w", err)
	}
	return nil
}
