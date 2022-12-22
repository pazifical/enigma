package encryption

import "github.com/TwoWaySix/enigma/internal"

type Reader struct {
	filepaths []string
	index     int
}

func NewReader(filepaths []string) Reader {
	return Reader{filepaths: filepaths}
}

func (r *Reader) ReadNext() (internal.UnencryptedFile, bool, error) {
	if r.index >= len(r.filepaths) {
		return internal.UnencryptedFile{}, false, nil
	}
	f, err := r.read(r.filepaths[r.index])
	if err != nil {
		// TODO: Implement
		return internal.UnencryptedFile{}, true, err
	}
	r.index += 1
	return f, true, nil
}

func (r *Reader) read(filePath string) (internal.UnencryptedFile, error) {
	// TODO: implement
	return internal.UnencryptedFile{}, nil
}
