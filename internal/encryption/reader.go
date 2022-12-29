package encryption

import (
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"os"
	"path/filepath"
)

type Reader struct {
	directoryPath string
	readFiles     chan internal.UnencryptedFile
}

func NewReader(directory string, readFiles chan internal.UnencryptedFile) Reader {
	return Reader{directoryPath: directory, readFiles: readFiles}
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
	err := filepath.Walk(r.directoryPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				//r.readFiles <- internal.UnencryptedFile{Data: nil} // TODO: Find an elegant solution
				return nil
			}
			return r.read(path)
		})
	if err != nil {
		return fmt.Errorf("finding all files in %s : %w", r.directoryPath, err)
	}
	// TODO: Close channel?
	close(r.readFiles)
	return nil
}

func (r *Reader) read(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("reading file %s : %w", filePath, err)
	}
	r.readFiles <- internal.UnencryptedFile{
		Data: data,
		Path: filePath,
	}
	return nil
}
