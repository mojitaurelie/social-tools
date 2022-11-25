package config

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed config.template.yml
var template []byte

var config Config

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
	log.Printf("INFO loading configuration file at %s", cpath)
	f, err := os.Open(cpath)
	if err != nil {
		return err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}
	log.Println("INFO configuration loaded!")
	return nil
}

func Current() Config {
	return config
}
