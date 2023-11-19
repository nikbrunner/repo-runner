package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func isSshUrl(url string) bool {
	return strings.Contains(url, "git@")
}

func parseGitUrl(gitUrl string) (string, string, error) {
	gitUrl = strings.TrimSuffix(gitUrl, ".git")

	var parts []string

	if isSshUrl(gitUrl) {
		parts = strings.Split(gitUrl, ":")
		if len(parts) < 2 {
			return "", "", fmt.Errorf("invalid Git Url")
		}

		parts = strings.Split(parts[1], "/")
	} else { // Assume HTTP/HTTPS format
		parts = strings.Split(gitUrl, "/")
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

func cloneRepo(gitUrl string) {
	config := loadConfig()

	reposBasePath := expandPath(config.ReposBasePath)
	if reposBasePath == "" {
		printNegative("Error getting base path", nil)
		return
	}

	separator := config.Separator
	if separator == "" {
		separator = "@"
	}

	// Parsing the Git URL to get username and repo name
	username, repoName, err := parseGitUrl(gitUrl)
	if err != nil {
		printNegative("Error parsing Git URL", err)
	}

	fullPath := fmt.Sprintf("%s/%s%s%s", reposBasePath, username, separator, repoName)

	if directoryExists(fullPath) {
		printNegative("Directory already exists!", nil)
		return
	}

	printPositive("Cloning repostitory to " + fullPath)

	cmd := exec.Command("git", "clone", gitUrl, fullPath)
	cmd.Stderr = os.Stderr // This will print any error output from the git command
	err = cmd.Run()
	if err != nil {
		printNegative("Error cloning repository:", err)
		return
	}

	printPositive("Repository cloned successfully")
}
