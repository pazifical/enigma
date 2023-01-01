package internal

import (
	"os"
	"strings"
	"testing"
)

var existingEncryptionInputDirectory = "./testdata/existing"
var nonExistingEncryptionInputDirectory = "./testdata/nonExisting"
var existingEncryptionOutputFile = "./testdata/existing.roll"
var nonExistingEncryptionOutputFile = "./testdata/nonExisting.roll"

var existingDecryptionOutputDirectory = "./testdata/decrypted/existing"
var nonExistingDecryptionOutputDirectory = "./testdata/decrypted/nonExisting"
var existingDecryptionInputFile = existingEncryptionOutputFile
var nonExistingDecryptionInputFile = nonExistingEncryptionOutputFile

var validKey = strings.Repeat("0", 16)
var inValidKey = strings.Repeat("0", 1)

func TestMain(m *testing.M) {
	err := os.MkdirAll(existingEncryptionInputDirectory, 0755)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(existingEncryptionOutputFile)
	if err != nil {
		return
	}
	f.Close()
	err = os.MkdirAll(existingDecryptionOutputDirectory, 0755)
	if err != nil {
		panic(err)
	}

	m.Run()

	err = os.Remove(existingEncryptionOutputFile)
	if err != nil {
		panic(err)
	}
	err = os.Remove(existingEncryptionInputDirectory)
	if err != nil {
		panic(err)
	}
	err = os.Remove(existingDecryptionOutputDirectory)
	if err != nil {
		panic(err)
	}

}

func TestCorrectEncryptionConfig(t *testing.T) {
	config := Config{
		Mode:       "roll",
		InputPath:  existingEncryptionInputDirectory,
		OutputPath: nonExistingEncryptionOutputFile,
		Key:        validKey,
	}
	err := config.validate()
	if err != nil {
		t.Errorf("should not occur: %v", err)
	}
}

func TestCorrectDecryptionConfig(t *testing.T) {
	config := Config{
		Mode:       "unroll",
		InputPath:  existingDecryptionInputFile,
		OutputPath: existingDecryptionOutputDirectory,
		Key:        validKey,
	}
	err := config.validate()
	if err != nil {
		t.Errorf("should not occur: %v", err)
	}
}

func TestValidKey(t *testing.T) {
	err := validateAESKey(validKey)
	if err != nil {
		t.Errorf("should not occur: %v", err)
	}
}

func TestInvalidKey(t *testing.T) {
	err := validateAESKey(inValidKey)
	if err == nil {
		t.Errorf("an error should be thrown with invalid encryption key: %s", inValidKey)
	}
}

func TestWrongEncryptionInputPath(t *testing.T) {
	config := Config{
		Mode:       "roll",
		InputPath:  nonExistingEncryptionInputDirectory,
		OutputPath: existingEncryptionOutputFile,
		Key:        validKey,
	}

	err := config.validate()
	if err == nil {
		t.Errorf("an error should be thrown with non existing input path: %s", nonExistingEncryptionInputDirectory)
	}
}

func TestWrongEncryptionOutputPath(t *testing.T) {
	config := Config{
		Mode:       "roll",
		InputPath:  existingEncryptionInputDirectory,
		OutputPath: existingEncryptionOutputFile,
		Key:        validKey,
	}

	err := config.validate()
	if err == nil {
		t.Errorf("an error should be thrown with existing output path: %s", existingEncryptionOutputFile)
	}
}

func TestWrongDecryptionInputPath(t *testing.T) {
	config := Config{
		Mode:       "unroll",
		InputPath:  nonExistingDecryptionInputFile,
		OutputPath: existingDecryptionOutputDirectory,
		Key:        validKey,
	}

	err := config.validate()
	if err == nil {
		t.Errorf("an error should be thrown with non existing input path: %s", nonExistingEncryptionInputDirectory)
	}
}

func TestWrongDecryptionOutputPath(t *testing.T) {
	config := Config{
		Mode:       "unroll",
		InputPath:  existingDecryptionInputFile,
		OutputPath: nonExistingDecryptionOutputDirectory,
		Key:        validKey,
	}

	err := config.validate()
	if err == nil {
		t.Errorf("an error should be thrown with wrong output path: %s", nonExistingEncryptionInputDirectory)
	}
}
