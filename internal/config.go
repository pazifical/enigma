package internal

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Mode       string
	InputPath  string
	OutputPath string
	Key        string
}

func NewConfig(mode, inputPath, outputPath, key string) (Config, error) {
	return Config{
		Mode:       mode,
		InputPath:  inputPath,
		OutputPath: outputPath,
		Key:        key,
	}, nil
}

func (c *Config) validate() error {
	err := validateAESKey(c.Key)
	if err != nil {
		return fmt.Errorf("validating config : %w", err)
	}
	switch c.Mode {
	case "roll":
		err := c.validateEncryptionValues()
		if err != nil {
			return fmt.Errorf("validating config : %w", err)
		}
	case "unroll":
		err := c.validateDecryptionValues()
		if err != nil {
			return fmt.Errorf("validating config : %w", err)
		}
	default:
		return fmt.Errorf("'%s' is not a valid mode", c.Mode)
	}
	return nil
}

func (c *Config) validateEncryptionValues() error {
	err := validateInputDirectory(c.InputPath)
	if err != nil {
		return err
	}
	err = validateOutputFile(c.OutputPath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) validateDecryptionValues() error {
	err := validateOutputDirectory(c.OutputPath)
	if err != nil {
		return err
	}
	err = validateInputFile(c.InputPath)
	if err != nil {
		return err
	}
	return nil
}

func validateInputFile(encryptedFilePath string) error {
	stat, err := os.Stat(encryptedFilePath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("input file does not exist: %s", encryptedFilePath)
	} else if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("input file is a directory: %s", encryptedFilePath)
	}
	return nil
}

func validateOutputDirectory(directory string) error {
	stat, err := os.Stat(directory)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		fileInfos, err := os.ReadDir(directory)
		if err != nil {
			return err
		}
		if len(fileInfos) > 0 {
			return fmt.Errorf("output directory '%s' is not empty", directory)
		}
	}
	// TODO: Check if is writable
	return nil
}

func validateAESKey(key string) error {
	if len(key) == 16 || len(key) == 24 || len(key) == 32 {
		return nil
	}
	return fmt.Errorf("given AES key is not of length 16, 24 or 32, but of length %d", len(key))
}

func validateInputDirectory(directory string) error {
	stat, err := os.Stat(directory)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("input path is not a directory")
	}
	return nil
}

func validateOutputFile(filePath string) error {
	stat, err := os.Stat(filePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if stat != nil {
		return fmt.Errorf("output file already exists and won't be overwritten: %s", filePath)
	}

	outputDirectory := filepath.Dir(filePath)
	dirStat, err := os.Stat(outputDirectory)
	if os.IsNotExist(err) {
		return fmt.Errorf("output directory does not exist: %s", outputDirectory)
	} else if err != nil {
		return err
	}
	if !dirStat.IsDir() {
		return fmt.Errorf("output directory is not a directory")
	}
	return nil
}

func NewConfigFromFlags() (Config, error) {
	rollCmd := flag.NewFlagSet("roll", flag.ExitOnError)
	rollInput := rollCmd.String("input", "", "TODO")
	rollOutput := rollCmd.String("output", "", "TODO")
	rollKey := rollCmd.String("key", "", "TODO")

	unrollCmd := flag.NewFlagSet("unroll", flag.ExitOnError)
	unrollInput := unrollCmd.String("input", "", "TODO")
	unrollOutput := unrollCmd.String("output", "", "TODO")
	unrollKey := unrollCmd.String("key", "", "TODO")

	if len(os.Args) < 2 {
		fmt.Println("expected 'roll' or 'unroll' subcommands")
		os.Exit(1)
	}

	var config Config
	switch os.Args[1] {
	case "roll":
		err := rollCmd.Parse(os.Args[2:])
		if err != nil {
			return Config{}, fmt.Errorf("parsing cli arguments")
		}
		config, err = NewConfig(os.Args[1], *rollInput, *rollOutput, *rollKey)
		if err != nil {
			return Config{}, fmt.Errorf("parsing cli arguments")
		}
	case "unroll":
		err := unrollCmd.Parse(os.Args[2:])
		if err != nil {
			return Config{}, fmt.Errorf("parsing cli arguments")
		}
		config, err = NewConfig(os.Args[1], *unrollInput, *unrollOutput, *unrollKey)
		if err != nil {
			return Config{}, fmt.Errorf("parsing cli arguments")
		}
	}
	return config, config.validate()
}
