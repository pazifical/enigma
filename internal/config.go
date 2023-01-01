package internal

import (
	"flag"
	"fmt"
	"os"
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
	// TODO: implement
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

	switch os.Args[1] {
	case "roll":
		err := rollCmd.Parse(os.Args[2:])
		if err != nil {
			return Config{}, fmt.Errorf("parsing cli arguments")
		}
		return NewConfig(os.Args[1], *rollInput, *rollOutput, *rollKey)
	case "unroll":
		err := unrollCmd.Parse(os.Args[2:])
		if err != nil {
			return Config{}, fmt.Errorf("parsing cli arguments")
		}
		return NewConfig(os.Args[1], *unrollInput, *unrollOutput, *unrollKey)
	default:
		return Config{}, fmt.Errorf("first argument has to be 'roll' or 'unroll'")
	}
}
