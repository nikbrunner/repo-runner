package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func parseGitUrl(gitUrl string) (string, string, error) {
	gitUrl = strings.TrimSuffix(gitUrl, ".git")

	var parts []string

	if strings.Contains(gitUrl, "git@") { // SSH format
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

func cloneRepo(gitUrl string) {
	config := loadConfig()

	config.RepoPath = expandPath(config.RepoPath)
	if config.RepoPath == "" {
		return
	}

	separator := config.Separator
	if separator == "" {
		separator = "@"
	}

	// Parsing the Git URL to get username and repo name
	username, repoName, err := parseGitUrl(gitUrl)
	if err != nil {
		printNegative("Error parsing Git URL", &err)
	}

	fullPath := fmt.Sprintf("%s/%s%s%s", config.RepoPath, username, separator, repoName)

	printPositive("Cloning repostitory to " + fullPath)

	cmd := exec.Command("git", "clone", gitUrl, fullPath)
	err = cmd.Run()
	if err != nil {
		printNegative("Error cloning repository:", &err)
		return
	}

	printPositive("Repository cloned successfully")
}
