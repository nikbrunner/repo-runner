package main

import (
	"os"
	"strings"
)

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func readDirs(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs, nil
}

func isSshUrl(url string) bool {
	return strings.Contains(url, "git@")
}

func isHttpUrl(url string) bool {
	return strings.Contains(url, "http://") || strings.Contains(url, "https://")
}

func isValidGitUrl(url string) bool {
	return isSshUrl(url) || isHttpUrl(url)
}

func sanitizePath(path string) string {
	return strings.TrimSuffix(path, "/")
}
