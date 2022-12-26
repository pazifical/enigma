package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/TwoWaySix/enigma/deprecated"
	"github.com/TwoWaySix/enigma/internal"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Refactor to use non deprecated workflow
func main() {
	fmt.Println("ENIGMA")
	config := parseFlags()

	log.Printf("INFO: please enter encryption key:")
	_, err := fmt.Scanf("%s", &config.Key)
	if err != nil {
		log.Printf("ERROR: entered key is invalid: %v", err)
		os.Exit(-1)
	}

	err = deprecated.EncryptAll(config)
	if err != nil {
		log.Printf("ERROR: encryption failed: %v", err)
		os.Exit(-1)
	}

	err = deprecated.CreateTarFromRolls(config)
	if err != nil {
		log.Printf("ERROR: packaging failed: %v", err)
		os.Exit(-1)
	}
}

func parseFlags() internal.Config {
	var mode string
	var paths string
	var outPath string
	flag.StringVar(&mode, "mode", "unroll", "roll or unroll")
	flag.StringVar(&paths, "paths", "", "comma separated paths to files or directories")
	flag.StringVar(&outPath, "out", "./enigma.roll", "output file")
	flag.Parse()

	var config internal.Config
	switch mode {
	case "roll":
		config.Mode = mode
	case "unroll":
		config.Mode = mode
	default:
		log.Println("ERROR: mode has to be either 'roll' or 'unroll'")
		os.Exit(-1)
	}

	parts := strings.Split(paths, ",")
	for _, p := range parts {
		_, err := os.Stat(p)
		if errors.Is(os.ErrNotExist, err) {
			log.Printf("WARNING: path does not exist and will be ignored: %s", p)
		} else if err != nil {
			log.Printf("WARNING: path is not valid and will be ignored: %s", p)
		} else {
			config.Paths = append(config.Paths, p)
		}
	}

	if len(config.Paths) == 0 {
		if config.Mode == "roll" {
			log.Printf("INFO: no paths specified. Files in the current and subdirectories will be encrypted.")
			paths, err := findAllFiles(".")
			if err != nil {
				log.Printf("ERROR: looking for files current directory: %v", err)
				os.Exit(-1)
			}
			config.Paths = paths
		} else {
			log.Printf("INFO: no paths specified. Looking for .roll files to decrypt.")
			// TODO: Implement
		}
	}

	config.OutputPath = outPath

	return config
}

func findAllFiles(directoryPath string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(directoryPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if filepath.Base(path) == os.Args[0] {
				return nil
			}
			fmt.Println(path, info.Size())
			filePaths = append(filePaths, path)
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("finding all files in %s : %w", directoryPath, err)
	}
	return filePaths, nil
}
