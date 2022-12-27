package main

import (
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"github.com/TwoWaySix/enigma/internal/decryption"
	"github.com/TwoWaySix/enigma/internal/encryption"
	"log"
	"os"
)

// TODO: Refactor to use non deprecated workflow
func main() {
	fmt.Println("ENIGMA")
	config, err := internal.NewConfigFromFlags()
	if err != nil {
		log.Printf("ERROR: %v", err)
		os.Exit(-1)
	}

	switch config.Mode {
	case "roll":
		err = startRollJob(config)
	case "unroll":
		err = startUnrollJob(config)
	}
	if err != nil {
		log.Printf("ERROR: %v", err)
		os.Exit(-1)
	}
}

func startRollJob(config internal.Config) error {
	job, err := encryption.NewJob(config)
	if err != nil {
		return fmt.Errorf("starting roll job: %w", err)
	}
	err = job.Start()
	if err != nil {
		return fmt.Errorf("starting roll job: %w", err)
	}
	return nil
}

func startUnrollJob(config internal.Config) error {
	job, err := decryption.NewJob(config)
	if err != nil {
		return fmt.Errorf("starting unroll job: %w", err)
	}
	err = job.Start()
	if err != nil {
		return fmt.Errorf("starting unroll job: %w", err)
	}
	return nil
}
