package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type Enigma struct {
	gcm cipher.AEAD
}

func NewEnigma(key string) (Enigma, error) {
	key, err := validateEncryptionKey(key)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return Enigma{}, fmt.Errorf("initializing cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return Enigma{}, fmt.Errorf("initializing cipher GCM: %w", err)
	}

	return Enigma{
		gcm: gcm,
	}, nil
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
