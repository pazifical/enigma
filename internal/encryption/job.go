package encryption

import (
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"github.com/TwoWaySix/enigma/pkg/enigma"
	"log"
)

type Job struct {
	reader Reader
	// readFiles chan internal.UnencryptedFile
	processor enigma.Machine
	writer    Writer
}

func NewJob(config internal.Config) (Job, error) {
	readFiles := make(chan internal.UnencryptedFile)
	reader := NewReader(config.InputPath, readFiles)

	processor, err := enigma.NewMachine(config.Key)
	if err != nil {
		return Job{}, fmt.Errorf("creating new job: %w", err)
	}

	writer, err := NewWriter(config.OutputPath)
	if err != nil {
		return Job{}, fmt.Errorf("creating new job: %w", err)
	}

	return Job{
		reader:    reader,
		processor: processor,
		writer:    writer,
		//readFiles: readFiles,
	}, nil
}

func (j *Job) Start() error {
	err := j.reader.Start() // TODO: Error handling
	if err != nil {
		return fmt.Errorf("running job: %w", err)
	}

	for {
		input := <-j.reader.readFiles
		if input.Data == nil { // TODO: Find an elegant solution
			break
		}

		encrypted, err := j.processor.Encrypt(input)
		if err != nil {
			log.Printf("ERROR: running job: %v", err)
			continue
		}

		err = j.writer.Write(encrypted)
		if err != nil {
			log.Printf("ERROR: running job: %v", err)
			continue
		}
	}

	err = j.Finish()
	if err != nil {
		return fmt.Errorf("running job: %w", err)
	}
	return nil
}

func (j *Job) Finish() error {
	err := j.writer.Close()
	if err != nil {
		return fmt.Errorf("finishing job: %w", err)
	}
	return nil
}
