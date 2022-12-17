package internal

import "testing"

func TestFileDecryptionCausesNoErrors(t *testing.T) {
	enigma, err := NewEnigma("key")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
	err = enigma.DecryptFile("./testdata/asdf.txt.roll")
	if err != nil {
		t.Errorf("encountered an error: %v", err)
	}
}
