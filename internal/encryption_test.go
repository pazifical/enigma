package internal

import "testing"

func TestFileEncryption(t *testing.T) {
	encrypter, err := NewEncrypter("key")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	err = encrypter.EncryptFile("./testdata/asdf.txt")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
}

func TestDirectoryEncryption(t *testing.T) {
	encrypter, err := NewEncrypter("key")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	err = encrypter.EncryptDirectory("./testdata/testdirectory")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
}
