package internal

import (
	"archive/tar"
	"errors"
	"fmt"
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

func createTar(directoryPath string, tarPath string) error {
	f, err := os.Create(tarPath)
	if err != nil {
		return fmt.Errorf("creating tar from directory %s : %w", directoryPath, err)
	}
	defer f.Close()

	tw := tar.NewWriter(f)
	err = filepath.Walk(directoryPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			fmt.Println(path, info.Size())

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			hdr := &tar.Header{
				Name: path,
				Mode: 0600,
				Size: info.Size(),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				log.Fatal(err)
			}
			if _, err := tw.Write(content); err != nil {
				log.Fatal(err)
			}
			return nil
		})
	if err != nil {
		return fmt.Errorf("creating tar from directory %s : %w", directoryPath, err)
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("creating tar from directory %s : %w", directoryPath, err)
	}
	return nil
}
