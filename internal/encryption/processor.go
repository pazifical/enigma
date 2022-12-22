package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"io"
)

type Processor struct {
	gcm cipher.AEAD
}

func NewProcessor(encryptionKey string) (Processor, error) {
	key, err := validateAESEncryptionKey(encryptionKey)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return Processor{}, fmt.Errorf("initializing cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return Processor{}, fmt.Errorf("initializing cipher GCM: %w", err)
	}
	return Processor{gcm: gcm}, nil
}

func (p *Processor) Encrypt(file internal.UnencryptedFile) (internal.EncryptedFile, error) {
	nonce := make([]byte, p.gcm.NonceSize())
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return internal.EncryptedFile{}, fmt.Errorf("initializing nonce: %w", err)
	}
	return internal.EncryptedFile{
		Data: p.gcm.Seal(nonce, nonce, file.Data, nil),
		Path: file.Path,
	}, nil
}

func validateAESEncryptionKey(key string) (string, error) {
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
