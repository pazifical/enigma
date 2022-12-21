package internal

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CreateTarFromRolls(config Config) error {
	f, err := os.Create(config.OutPath)
	if err != nil {
		return fmt.Errorf("creating tar %s : %w", config.OutPath, err)
	}
	defer f.Close()

	tw := tar.NewWriter(f)

	for _, filePath := range config.Paths {
		encryptedFilePath := fmt.Sprintf("%s.roll", filePath)

		stat, err := os.Stat(encryptedFilePath)
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("WARNING: writing file to tar %s : %v", encryptedFilePath, err)
			continue
		} else if err != nil {
			log.Printf("WARNING: writing file to tar %s : %v", encryptedFilePath, err)
			continue
		}

		content, err := os.ReadFile(encryptedFilePath)
		if err != nil {
			log.Printf("WARNING: writing file to tar %s : %v", encryptedFilePath, err)
			continue
		}

		hdr := &tar.Header{
			Name: encryptedFilePath,
			Mode: 0600,
			Size: stat.Size(),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Printf("WARNING: writing file to tar %s : %v", encryptedFilePath, err)
			continue
		}
		if _, err := tw.Write(content); err != nil {
			log.Printf("WARNING: writing file to tar %s : %v", encryptedFilePath, err)
			continue
		}
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("creating tar %s : %w", config.OutPath, err)
	}
	return nil
}

func UntarAll(config Config) {
	for _, p := range config.Paths {
		err := UntarRoll(p, config.OutPath)
		if err != nil {
			log.Printf("ERROR: unpacking all tars : %v", err)
		}
	}
}

func UntarRoll(tarPath string, outPath string) error {
	stat, err := os.Stat(outPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(outPath, 0750)
		if err != nil {
			return fmt.Errorf("unpacking to %s : %v", outPath, err)
		}
	} else if err != nil {
		return fmt.Errorf("unpacking to %s : %v", outPath, err)
	} else if !stat.IsDir() {
		return fmt.Errorf("%s has to be a directory", outPath)
	}

	tarFile, err := os.Open(tarPath)
	if err != nil {
		return err
	}

	tr := tar.NewReader(tarFile)
	for {
		header, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		target := filepath.Join(outPath, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			err = os.MkdirAll(filepath.Dir(target), 0750)
			if !errors.Is(os.ErrExist, err) && err != nil {
				return err
			}
			f, err := os.Create(target)
			// f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}
			f.Close()
		}
	}
}
