package enigma

import (
	"github.com/TwoWaySix/enigma/internal"
	"strings"
	"testing"
)

var key = strings.Repeat("0", 16)

func TestByteEncryptionAndDecryption(t *testing.T) {
	text := "hello"
	textBytes := []byte(text)
	unencrypted := internal.UnencryptedFile{Data: textBytes}

	processor, err := NewMachine(key)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	encrypted, err := processor.Encrypt(unencrypted)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	decrypted, err := processor.Decrypt(encrypted)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	if text != string(decrypted.Data) {
		t.Errorf("%s != %s", text, string(decrypted.Data))
	}
}
