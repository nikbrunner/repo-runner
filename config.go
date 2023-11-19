package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func loadConfig() Config {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)

	if err != nil {
		fmt.Println("Error decoding config file: ", err)
		os.Exit(1)
	}

	return config
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
