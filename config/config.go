package config

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"
)

//go:embed config.template.yml
var template []byte

func Load(cpath string) error {
	if len(cpath) == 0 {
		return errors.New("the configuration file path cannot be empty")
	}
	if _, err := os.Stat(cpath); err != nil {
		log.Printf("INFO no configuration found at %s, writing a new one", cpath)
		err = os.WriteFile(cpath, template, 0644)
		if err != nil {
			return fmt.Errorf("an error occured while writing a new configuration file: %s", err)
		}
	}

	return nil
}
