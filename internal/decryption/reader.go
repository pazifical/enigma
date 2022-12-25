package decryption

import "github.com/TwoWaySix/enigma/internal"

type Reader struct {
	zipPath string
}

func NewReader(zipPath string) (Reader, error) {
	// TODO: implement
	return Reader{zipPath: zipPath}, nil
}

func (r *Reader) ReadNext() (internal.EncryptedFile, bool, error) {
	// TODO: implement
	return internal.EncryptedFile{}, true, nil
}

func (r *Reader) Close() error {
	// TODO: implement
	return nil
}
