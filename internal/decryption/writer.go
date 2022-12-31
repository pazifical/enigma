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
	newFilePath := filepath.Join(w.outDirectory, decrypted.Path)
	f, err := os.Create(newFilePath)
	if err != nil {
		return fmt.Errorf("writing file '%s' from zip: %w", newFilePath, err)
	}
	_, err = f.Write(decrypted.Data)
	if err != nil {
		return fmt.Errorf("writing file '%s' from zip: %w", newFilePath, err)
	}
	return nil
}
