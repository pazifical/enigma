package test

import (
	"github.com/TwoWaySix/enigma/internal"
	"github.com/TwoWaySix/enigma/internal/decryption"
	"github.com/TwoWaySix/enigma/internal/encryption"
	"os"
	"path/filepath"
	"testing"
)

var inputDir = "./testdata/input"
var inputFile1Path = filepath.Join(inputDir, "file1.txt")
var encryptionDir = "./testdata/encrypted"
var encryptedFilePath = filepath.Join(encryptionDir, "file.zip")
var decryptionDir = "./testdata/decrypted"
var decryptionFilePath = filepath.Join(decryptionDir, "file1.txt")

func TestMain(m *testing.M) {
	err := os.MkdirAll(inputDir, 0755)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(encryptionDir, 0755)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(decryptionDir, 0755)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(inputFile1Path)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString("ENIGMA")
	if err != nil {
		panic(err)
	}

	m.Run()

	err = os.Remove(inputFile1Path)
	if err != nil {
		panic(err)
	}
	err = os.Remove(inputDir)
	if err != nil {
		panic(err)
	}
}

func TestEncryptionDecryption(t *testing.T) {
	// ENCRYPTION
	config := internal.Config{
		Mode:       "roll",
		InputPath:  inputDir,
		OutputPath: encryptedFilePath,
		Key:        "asdf",
	}

	encryptionJob, err := encryption.NewJob(config)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = encryptionJob.Start()
	if err != nil {
		t.Errorf(err.Error())
	}

	// DECRYPTION
	config.Mode = "unroll"
	config.InputPath = encryptedFilePath
	config.OutputPath = decryptionDir

	decryptionJob, err := decryption.NewJob(config)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = decryptionJob.Start()
	if err != nil {
		t.Errorf(err.Error())
	}
}
