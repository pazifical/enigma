package decryption

import (
	"archive/zip"
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"log"
)

type Reader struct {
	zipPath   string
	archive   *zip.ReadCloser
	readFiles chan internal.EncryptedFile
}

func NewReader(zipPath string, readFiles chan internal.EncryptedFile) (Reader, error) {
	archive, err := zip.OpenReader(zipPath)
	if err != nil {
		return Reader{}, fmt.Errorf("creating new reader: %w", err)
	}

	return Reader{
		zipPath:   zipPath,
		archive:   archive,
		readFiles: readFiles,
	}, nil
}

func (r *Reader) Start() error {
	var err error
	go func() {
		err = r.readAllFiles()
		if err != nil {
			return
		}
	}()
	return err
}

func (r *Reader) readAllFiles() error {
	archive, err := zip.OpenReader(r.zipPath)
	if err != nil {
		return fmt.Errorf("reading all files: %w", err)
	}
	for _, f := range archive.File {
		err := r.read(f)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}
	}
	close(r.readFiles)
	return nil
}

func (r *Reader) read(f *zip.File) error {
	readCloser, err := f.Open()
	if err != nil {
		return fmt.Errorf("reading %s : %w", f.Name, err)
	}

	if f.FileInfo().IsDir() {
		return nil
	}

	bytes := f.FileInfo().Size()
	data := make([]byte, bytes, bytes)
	_, err = readCloser.Read(data)
	if err != nil {
		// TODO: handle properly
		// return fmt.Errorf("reading %s : %w", f.Name, err)
	}
	r.readFiles <- internal.EncryptedFile{Data: data, Path: f.Name} // TODO: Check if f.Name is really the correct path
	return nil
}

func (r *Reader) Close() error {
	err := r.archive.Close()
	if err != nil {
		return fmt.Errorf("closing archive: %w", err)
	}
	return nil
}
