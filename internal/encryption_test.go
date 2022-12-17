package internal

import "testing"

func TestFileEncryptionCausesNoErrors(t *testing.T) {
	encrypter, err := NewEnigma("key")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	err = encrypter.EncryptFile("./testdata/asdf.txt")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
}

func TestDirectoryEncryptionCausesNoErrors(t *testing.T) {
	encrypter, err := NewEnigma("key")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	err = encrypter.EncryptDirectory("./testdata/testdirectory")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
}
