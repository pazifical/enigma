package main

import (
	"bufio"
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"github.com/TwoWaySix/enigma/internal/decryption"
	"github.com/TwoWaySix/enigma/internal/encryption"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting ENIGMA")
	var err error
	if len(os.Args) > 1 {
		err = normalMode()
	} else {
		err = simpleDecryptMode()
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}

func simpleDecryptMode() error {
	fmt.Println("Starting simple decrypt mode.")
	fmt.Println("Looking for files with '.roll' extension.")
	dirEntries, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	for _, entry := range dirEntries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".roll") {
			continue
		}
		err := simpleDecryptFile(entry)
		if err != nil {
			fmt.Printf("ERROR: %v", err)
			continue
		}
	}
	return nil
}

func simpleDecryptFile(entry os.DirEntry) error {
	fmt.Printf("Decrypting file %s\n", entry.Name())
	fmt.Printf("Please enter encryption key: ")
	reader := bufio.NewReader(os.Stdin)
	key, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("decrypting file '%s': %w", entry.Name(), err)
	}

	config := internal.Config{
		Mode:       "unroll",
		InputPath:  entry.Name(),
		OutputPath: ".",
		Key:        key,
	}

	err = startUnrollJob(config)
	if err != nil {
		return fmt.Errorf("decrypting file '%s': %w", entry.Name(), err)
	}
	return nil
}

func normalMode() error {
	config, err := internal.NewConfigFromFlags()
	if err != nil {
		return err
	}

	switch config.Mode {
	case "roll":
		err = startRollJob(config)
	case "unroll":
		err = startUnrollJob(config)
	}
	if err != nil {
		return err
	}
	return nil
}

func startRollJob(config internal.Config) error {
	job, err := encryption.NewJob(config)
	if err != nil {
		return fmt.Errorf("roll job: %w", err)
	}
	err = job.Start()
	if err != nil {
		return fmt.Errorf("roll job: %w", err)
	}
	return nil
}

func startUnrollJob(config internal.Config) error {
	job, err := decryption.NewJob(config)
	if err != nil {
		return fmt.Errorf("unroll job: %w", err)
	}
	err = job.Start()
	if err != nil {
		return fmt.Errorf("unroll job: %w", err)
	}
	return nil
}
