package deprecated

import (
	"errors"
	"github.com/TwoWaySix/enigma/internal"
	"os"
	"path/filepath"
	"testing"
)

func TestPackingAndUnpacking(t *testing.T) {
	inputDirectory := "./testdata/one"
	outputDirectory := "./testdata/two"
	fileName := "test.txt"
	encryptedFileName := fileName + ".roll"

	filePath := filepath.Join(inputDirectory, fileName)
	encryptedFilePath := filepath.Join(inputDirectory, encryptedFileName)
	tarPath := filepath.Join(inputDirectory, fileName+".tar")

	err := os.MkdirAll(inputDirectory, 0750)
	if err != nil && errors.Is(os.ErrExist, err) {
		t.Errorf("Cannot create %s: %v", inputDirectory, err)
	}
	f, err := os.Create(encryptedFilePath)
	if err != nil {
		t.Errorf("Cannot create %s: %v", filePath, err)
	}
	f.Close()

	config := internal.Config{
		Mode:    "roll",
		Paths:   []string{filePath},
		Key:     "asdf",
		OutPath: tarPath,
	}
	err = CreateTarFromRolls(config)
	if err != nil {
		t.Errorf("Cannot create tar %s: %v", tarPath, err)
	}

	_, err = os.Stat(tarPath)
	if err != nil {
		t.Errorf("Cannot stat tar %s: %v", tarPath, err)
	}

	config = internal.Config{
		Mode:    "unroll",
		Paths:   []string{tarPath},
		Key:     "asdf",
		OutPath: outputDirectory,
	}
	err = UntarRoll(tarPath, outputDirectory)
	if err != nil {
		t.Errorf("Cannot untar %s: %v", tarPath, err)
	}

	_, err = os.Stat(filepath.Join(outputDirectory, encryptedFilePath))
	if err != nil {
		t.Errorf("Cannot untar %s: %v", tarPath, err)
	}

	err = os.RemoveAll(inputDirectory)
	if err != nil {
		t.Errorf("Cannot delete %s: %v", inputDirectory, err)
	}
	err = os.RemoveAll(outputDirectory)
	if err != nil {
		t.Errorf("Cannot delete %s: %v", outputDirectory, err)
	}
}
