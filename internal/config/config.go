package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Key string `json:"key"`
}

// LoadConfig loads a configuration file from the given path
func LoadConfig(path string) (*Config, error) {

	cfg, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Add logic to parse the configuration file
	var config Config
	err = json.Unmarshal(cfg, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &config, nil
}
