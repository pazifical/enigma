package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/TwoWaySix/enigma/internal"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("ENIGMA")
	config := parseFlags()

	log.Printf("INFO: please enter encryption key:")
	_, err := fmt.Scanf("%s", &config.Key)
	if err != nil {
		log.Printf("ERROR: entered key is invalid.")
		os.Exit(-1)
	}

	internal.EncryptAll(config)

}

func parseFlags() internal.Config {
	var mode string
	var paths string
	flag.StringVar(&mode, "mode", "unroll", "roll or unroll")
	flag.StringVar(&paths, "paths", "", "comma separated paths to files or directories")
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
			log.Printf("INFO: no paths specified. Files and directories in the current directory will be encrypted.")
			files, err := os.ReadDir(".")
			if err != nil {
				log.Printf("ERROR: cannot read current directory content: %v", err)
				os.Exit(-1)
			}
			for _, f := range files {
				if f.Name() != os.Args[0] {
					config.Paths = append(config.Paths, f.Name())
				}
			}
		} else {
			log.Printf("INFO: no paths specified. Looking for .roll files to decrypt.")
			// TODO: Implement
		}
	}
	return config
}
