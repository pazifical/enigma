package decryption

import (
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"os"
	"path/filepath"
)

type Writer struct {
	outDirectory string
}

func NewWriter(outDirectory string) (Writer, error) {
	// TODO: Only continue if outDirectory is empty!
	return Writer{outDirectory: outDirectory}, nil
}

func (w *Writer) Write(decrypted internal.UnencryptedFile) error {
	newPath := filepath.Join(w.outDirectory, decrypted.Path)
	f, err := os.Create(newPath)
	if err != nil {
		return fmt.Errorf("writing file '%s' from zip: %w", newPath, err)
	}
	_, err = f.Write(decrypted.Data)
	if err != nil {
		return fmt.Errorf("writing file '%s' from zip: %w", newPath, err)
	}
	return nil
}
