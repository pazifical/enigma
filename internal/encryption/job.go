package encryption

import "github.com/TwoWaySix/enigma/internal"

type Job struct {
	reader    Reader
	processor Processor
	writer    Writer
}

func NewJob(config internal.Config) Job {
	// TODO: Validate
	return Job{
		reader:    NewReader(config.Paths),
		processor: NewProcessor(config.Key),
		writer:    NewWriter(config.OutPath),
	}
}

func (j *Job) Start() error {
	// TODO: Implement
	return nil
}
