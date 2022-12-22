package encryption

import (
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Mkdir("./testdata", 0755)

	m.Run()
}

func TestWriter(t *testing.T) {
	filePath := "./testdata/testfile.txt"
	file, err := os.Create(filePath)
	if err != nil {
		t.Errorf(err.Error())
	}
	file.WriteString("Hello")
	file.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf(err.Error())
	}

	outPath := "./testdata/packaged.zip"

	writer, err := NewWriter(outPath)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer writer.Close()

	err = writer.Write(internal.EncryptedFile{
		Data: data,
		Path: filePath,
	})
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
