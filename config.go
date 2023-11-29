package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const defaultConfigFileName = "defaultConfig.json"

//go:embed defaultConfig.json
var defaultConfigFile embed.FS

//go:embed layouts/default.sh
var defaultLayoutScript string

type LayoutType string

const (
	LayoutDefault LayoutType = "default"
)

type Config struct {
	ReposBasePath string     `json:"reposBasePath"`
	Layout        LayoutType `json:"layout"`
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

func loadDefaultConfig() (Config, error) {
	var defaultConfig Config

	// Read the embedded file
	configFileData, err := defaultConfigFile.ReadFile(defaultConfigFileName)
	if err != nil {
		return Config{}, fmt.Errorf("error reading embedded config file: %w", err)
	}

	// Decode the configuration from the embedded file data
	if err := json.Unmarshal(configFileData, &defaultConfig); err != nil {
		return Config{}, fmt.Errorf("error decoding embedded config file: %w", err)
	}

	return defaultConfig, nil
}

func validateConfig(config Config) error {
	if config.ReposBasePath == "" {
		return fmt.Errorf("reposBasePath is required")
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

func validateLayout(layout LayoutType) error {
	switch layout {
	case LayoutDefault:
		return nil
	default:
		return fmt.Errorf(fmt.Sprintf("invalid layout: '%s'", layout))
	}
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
