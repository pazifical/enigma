package decryption

import (
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"github.com/TwoWaySix/enigma/pkg/enigma"
)

type Job struct {
	reader    Reader
	processor enigma.Machine
	writer    Writer
}

func NewJob(config internal.Config) (Job, error) {
	readFiles := make(chan internal.EncryptedFile)
	reader, err := NewReader(config.InputPath, readFiles)
	if err != nil {
		return Job{}, fmt.Errorf("creating new job: %w", err)
	}

	processor, err := enigma.NewMachine(config.Key)
	if err != nil {
		return Job{}, fmt.Errorf("creating new job: %w", err)
	}

	writer, err := NewWriter(config.OutPath)
	if err != nil {
		return Job{}, fmt.Errorf("creating new job: %w", err)
	}

	return Job{
		reader:    reader,
		processor: processor,
		writer:    writer,
	}, nil
}

func (j *Job) Start() error {
	err := j.reader.Start()
	if err != nil {
		return fmt.Errorf("running job: %w", err)
	}

	for {
		encrypted := <-j.reader.readFiles
		if encrypted.Data == nil { // TODO: Find an elegant solution
			break
		}

		decrypted, err := j.processor.Decrypt(encrypted)
		if err != nil {
			return fmt.Errorf("running decryption job: %w", err)
		}

		err = j.writer.Write(decrypted)
		if err != nil {
			return fmt.Errorf("running decryption job: %w", err)
		}
	}

	err = j.Finish()
	if err != nil {
		return fmt.Errorf("running decryption job: %w", err)
	}
	return nil
}

func (j *Job) Finish() error {
	err := j.reader.Close()
	if err != nil {
		return fmt.Errorf("finalizing job: %w", err)
	}
	return nil
}
