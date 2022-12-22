package decryption

import (
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
)

type Job struct {
	reader    Reader
	processor internal.Processor
	writer    Writer
}

func NewJob(config internal.Config) (Job, error) {
	reader := NewReader(config.Paths)

	processor, err := internal.NewProcessor(config.Key)
	if err != nil {
		return Job{}, fmt.Errorf("creating new job: %w", err)
	}

	writer := NewWriter(config.OutPath)
	if err != nil {
		return Job{}, fmt.Errorf("creating new job: %w", err)
	}

	return Job{
		reader:    reader,
		processor: processor,
		writer:    writer,
	}, nil
}
