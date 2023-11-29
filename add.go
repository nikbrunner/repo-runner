package main

import (
	"fmt"
	"strings"
)

func addRepo(config Config, gitUrl string) {
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

func getClonePath(gitUrl string, config Config) string {
	if !isValidGitUrl(gitUrl) {
		printNegative("Invalid Git URL", nil)
		return ""
	}

	username, repoName, err := parseGitUrl(gitUrl)
	if err != nil {
		printNegative("Error parsing Git URL", err)
	}

	usernameDir := fmt.Sprintf("%s/%s", config.ReposBasePath, username)

	// Check if the repository for the user exists
	if directoryExists(usernameDir) {
		printPositive(fmt.Sprintf("Directory for GitHub user '%s' found!\nAdding new repository '%s'.", username, repoName))
	} else {
		printPositive(fmt.Sprintf("Directory for GitHub user '%s' not found!\nCreating directory.", username))
	}

	clonePath := fmt.Sprintf("%s/%s/%s", config.ReposBasePath, username, repoName)

	return clonePath
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
