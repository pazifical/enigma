package enigma

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"io"
)

type Machine struct {
	gcm cipher.AEAD
}

func NewMachine(encryptionKey string) (Machine, error) {
	key, err := validateAESEncryptionKey(encryptionKey)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return Machine{}, fmt.Errorf("initializing cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return Machine{}, fmt.Errorf("initializing cipher GCM: %w", err)
	}
	return Machine{gcm: gcm}, nil
}

func (p *Machine) Encrypt(file internal.UnencryptedFile) (internal.EncryptedFile, error) {
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

func (p *Machine) Decrypt(file internal.EncryptedFile) (internal.UnencryptedFile, error) {
	nonce := file.Data[:p.gcm.NonceSize()]
	data := file.Data[p.gcm.NonceSize():]
	decrypted, err := p.gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return internal.UnencryptedFile{}, fmt.Errorf("decrypting: %w", err)
	}
	return internal.UnencryptedFile{
		Data: decrypted,
		Path: file.Path,
	}, nil
}

func validateAESEncryptionKey(key string) (string, error) {
	if len(key) != 16 || len(key) != 24 || len(key) != 32 {
		return key, fmt.Errorf("given AES key is not of length 16, 24 or 32")
	}
	return key, nil
}
