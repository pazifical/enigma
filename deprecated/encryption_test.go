package deprecated

import (
	"testing"
)

func TestFileEncryptionCausesNoErrors(t *testing.T) {
	enigma, err := NewEnigma("key")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}

	err = enigma.EncryptFile("./testdata/asdf.txt", "./testdata/asdf.txt.roll")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
}
