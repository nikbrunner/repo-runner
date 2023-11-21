package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const defaultSeparator = "@"

func loadConfig() (Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("error decoding config file: %w", err)
	}

	// Validate and set defaults
	config.ReposBasePath = expandPath(config.ReposBasePath)
	if config.ReposBasePath == "" {
		return Config{}, fmt.Errorf("repository base path is not set")
	}

	if config.Separator == "" {
		config.Separator = defaultSeparator
	}

	return config, nil
}

func expandPath(path string) string {
	if strings.Contains(path, "$HOME") {
		home, err := os.UserHomeDir()
		if err != nil {
			printNegative("Error getting home directory:", err)
			return ""
		}
		return strings.Replace(path, "$HOME", home, 1)
	}
	return path
}
