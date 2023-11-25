package main

import (
	"fmt"
	"os"
	"strings"
)

func isSshUrl(url string) bool {
	return strings.Contains(url, "git@")
}

func isHttpUrl(url string) bool {
	return strings.Contains(url, "http://") || strings.Contains(url, "https://")
}

func isValidGitUrl(url string) bool {
	return isSshUrl(url) || isHttpUrl(url)
}

func parseGitUrl(gitUrl string) (string, string, error) {
	gitUrl = strings.TrimSuffix(gitUrl, ".git")

	var parts []string

	if isSshUrl(gitUrl) {
		parts = strings.Split(gitUrl, ":")
		parts = strings.Split(parts[1], "/")
	} else if isHttpUrl(gitUrl) {
		parts = strings.Split(gitUrl, "/")
	} else {
		return "", "", fmt.Errorf("invalid Git Url")
	}

	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid Git Url")
	}

	username := parts[len(parts)-2]
	repoName := parts[len(parts)-1]

	return username, repoName, nil
}

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func getClonePath(gitUrl string, config Config) string {
	if !isValidGitUrl(gitUrl) {
		printNegative("Invalid Git URL", nil)
		return ""
	}

	username, repoName, err := parseGitUrl(gitUrl)
	if err != nil {
		printNegative("Error parsing Git URL", err)
	}

	clonePath := fmt.Sprintf("%s/%s%s%s", config.ReposBasePath, username, config.Separator, repoName)

	return clonePath
}

func cloneRepo(config Config, gitUrl string) {
	clonePath := getClonePath(gitUrl, config)

	if directoryExists(clonePath) {
		printNegative("Directory already exists!", nil)
		return
	}

	printPositive("Cloning repository...")

	if err := git("clone", gitUrl, clonePath); err != nil {
		printNegative("Error cloning repository:", err)
		return
	}

	printPositive("Repository cloned successfully")
}
