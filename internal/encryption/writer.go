package encryption

import (
	"archive/zip"
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"os"
)

type Writer struct {
	filepath string
	file     *os.File
	writer   *zip.Writer
}

func NewWriter(filepath string) (Writer, error) {
	archive, err := os.Create(filepath)
	if err != nil {
		return Writer{}, fmt.Errorf("initializing writer: %w", err)
	}
	writer := zip.NewWriter(archive)
	return Writer{filepath: filepath, file: archive, writer: writer}, nil
}

func (w *Writer) Write(file internal.EncryptedFile) error {
	zipWriter, err := w.writer.Create(file.Path)
	if err != nil {
		return fmt.Errorf("writing '%s' to zip : %v", file.Path, err)
	}

	_, err = zipWriter.Write(file.Data)
	if err != nil {
		return fmt.Errorf("writing '%s' to zip : %v", file.Path, err)
	}

	return nil
}

func (w *Writer) Close() error {
	err := w.writer.Close()
	if err != nil {
		return fmt.Errorf("closing writer: %w", err)
	}
	err = w.file.Close()
	if err != nil {
		return fmt.Errorf("closing writer: %w", err)
	}
	return nil
}
