package deprecated

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Check, if it is a tar archive and then unpack it!
func (e *deprecated.Enigma) DecryptFile(filePath string) error {
	encrypted, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("decrypting file %s : %w", filePath, err)
	}
	newFilePath := strings.Replace(filePath, filepath.Ext(filePath), "", -1)
	f, err := os.Create(newFilePath)
	if err != nil {
		return fmt.Errorf("decrypting file %s : %w", filePath, err)
	}
	defer f.Close()

	decrypted, err := e.Decrypt(encrypted)
	if err != nil {
		return fmt.Errorf("decrypting file %s : %w", filePath, err)
	}

	_, err = f.Write(decrypted)
	if err != nil {
		return fmt.Errorf("decrypting file %s : %w", filePath, err)
	}

	return nil
}

func (e *deprecated.Enigma) Decrypt(data []byte) ([]byte, error) {
	nonce := data[:e.gcm.NonceSize()]
	data = data[e.gcm.NonceSize():]
	plainText, err := e.gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, fmt.Errorf("decrypting: %w", err)
	}
	return plainText, nil
}
