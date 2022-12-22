package encryption

import (
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"os"
)

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

	filePath := r.filepaths[r.index]
	f, err := r.read(filePath)
	if err != nil {
		return internal.UnencryptedFile{}, true, fmt.Errorf("reading next file %s : %w", filePath, err)
	}

	r.index += 1
	return f, true, nil
}

func (r *Reader) read(filePath string) (internal.UnencryptedFile, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return internal.UnencryptedFile{}, fmt.Errorf("reading file %s: %v", filePath, err)
	}
	return internal.UnencryptedFile{
		Data: content,
		Path: filePath,
	}, nil
}
