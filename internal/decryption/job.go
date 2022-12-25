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
	reader, err := NewReader(config.InputPath)
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
	for {
		encrypted, ok, err := j.reader.ReadNext()
		if err != nil {
			return fmt.Errorf("running decryption job: %w", err)
		}
		if !ok {
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

	err := j.Finish()
	if err != nil {
		return fmt.Errorf("running decryption job: %w", err)
	}
	return nil
}

func (j *Job) Finish() error {
	// TODO: implement
	return nil
}
