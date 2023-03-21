package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Source struct {
	Name    string `yaml:"name"`
	Source  string `yaml:"source"`
	Version string `yaml:"version"`
}

type Config struct {
	Target  string   `yaml:"target"`
	Output  string   `yaml:"output"`
	Sources []Source `yaml:"sources"`
}

func GetConfig(configFile string) Config {

	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("failed to read configuration in file %s due to error: %s", versed, err)
	}

	var config Config
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("failed to parse yaml file due to error: %s", err)
	}

	return config
}
