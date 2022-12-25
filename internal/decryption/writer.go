package decryption

import "github.com/TwoWaySix/enigma/internal"

type Writer struct {
	outDirectory string
}

func NewWriter(outDirectory string) (Writer, error) {
	// TODO: implement
	// TODO: Only continue if outDirectory is empty!
	return Writer{outDirectory: outDirectory}, nil
}

func (w Writer) Write(decrypted internal.UnencryptedFile) error {
	// TODO: implement
	return nil
}
