package internal

import (
	"os"
	"testing"
)

func TestEncryptionAndDecryption(t *testing.T) {
	text := "hello"
	textBytes := []byte(text)

	enigma, err := NewEnigma("key")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	encrypted, err := enigma.Encrypt(textBytes)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	decrypted, err := enigma.Decrypt(encrypted)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	if text != string(decrypted) {
		t.Errorf("%s != %s", text, string(decrypted))
	}
}

func TestFileEncryptionDecryption(t *testing.T) {
	filePath := "./testdata/dummy.txt"
	encryptedFilePath := "./testdata/dummy.txt.roll"
	content := "This is a lot more secure than an enigma."

	f, err := os.Create(filePath)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
	f.Close()

	enigma, err := NewEnigma("key")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	err = enigma.EncryptFile(filePath)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	err = enigma.DecryptFile(encryptedFilePath)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	f2, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
	newContent := string(f2)

	if content != newContent {
		t.Errorf("got %s, expected %s", newContent, content)
	}

	err = os.Remove(filePath)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
	err = os.Remove(encryptedFilePath)
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
}
