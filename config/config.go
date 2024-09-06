package config

import (
	"encoding/json"
	"os"
)

// DirectoryConfig holds the configuration for a single directory
type DirectoryConfig struct {
	Path      string `json:"path"`
	DeleteAll bool   `json:"delete_all"`
}

// Config holds the list of directory configurations
type Config struct {
	Directories []DirectoryConfig `json:"directories"`
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig(filename string) (*Config, error) {
	var cfg Config
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
