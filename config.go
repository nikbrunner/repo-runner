package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const defaultConfigFileName = "defaultConfig.json"

type LayoutType string

const (
	LayoutDefault LayoutType = "default"
)

type Config struct {
	ReposBasePath string     `json:"reposBasePath"`
	Separator     string     `json:"separator"`
	Layout        LayoutType `json:"layout"`
}

func validateLayout(layout LayoutType) error {
	switch layout {
	case LayoutDefault:
		return nil
	default:
		return fmt.Errorf(fmt.Sprintf("invalid layout: '%s'", layout))
	}
}

func loadDefaultConfig() (Config, error) {
	var defaultConfig Config

	defaultConfigFile, err := os.Open(defaultConfigFileName)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer defaultConfigFile.Close()

	decoder := json.NewDecoder(defaultConfigFile)
	if err := decoder.Decode(&defaultConfig); err != nil {
		return Config{}, fmt.Errorf("error decoding config file: %w", err)
	}

	return defaultConfig, nil
}

func validateConfig(config Config) error {
	if config.ReposBasePath == "" {
		return fmt.Errorf("reposBasePath is required")
	}

	if config.Separator == "" {
		return fmt.Errorf("separator is required")
	}

	if config.Layout == "" {
		return fmt.Errorf("layout is required")
	} else {
		if err := validateLayout(config.Layout); err != nil {
			return fmt.Errorf("error validating layout: %w", err)
		}
	}

	return nil
}

// Expand $HOME in the path
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

func loadConfig() (Config, error) {
	defaultConfig, defaultConfigErr := loadDefaultConfig()
	if defaultConfigErr != nil {
		return Config{}, fmt.Errorf("error loading default config: %w", defaultConfigErr)
	}

	validationErr := validateConfig(defaultConfig)
	if validationErr != nil {
		return Config{}, fmt.Errorf("error validating default config: %w", validationErr)
	}

	defaultConfig.ReposBasePath = expandPath(defaultConfig.ReposBasePath)

	return defaultConfig, nil
}
