package decryption

import "github.com/TwoWaySix/enigma/internal"

type Writer struct {
	outDirectory string
}

func NewWriter(outDirectory string) (Writer, error) {
	// TODO: implement
	// TODO: Only continue if outDirectory is empty!
	panic("not implemented")
	return Writer{outDirectory: outDirectory}, nil
}

func (w Writer) Write(decrypted internal.UnencryptedFile) error {
	// TODO: implement
	panic("not implemented")
	return nil
}
