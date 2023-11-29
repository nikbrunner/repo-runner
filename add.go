package main

import (
	"fmt"
	"strings"
)

func addRepo(config Config, gitUrl string) {
	clonePath := getClonePath(gitUrl, config)
	log := NewLogUtil()

	if directoryExists(clonePath) {
		log.Negative("Directory already exists!", nil)
		return
	}

	log.Positive("Cloning repository...")

	if err := git("clone", gitUrl, clonePath); err != nil {
		log.Negative("Error cloning repository:", err)
		return
	}

	log.Positive("Repository cloned successfully")
}

func getClonePath(gitUrl string, config Config) string {
	log := NewLogUtil()

	if !isValidGitUrl(gitUrl) {
		log.Negative("Invalid Git URL", nil)
		return ""
	}

	username, repoName, err := parseGitUrl(gitUrl)
	if err != nil {
		log.Negative("Error parsing Git URL", err)
	}

	usernameDir := fmt.Sprintf("%s/%s", config.ReposBasePath, username)

	// Check if the repository for the user exists
	if directoryExists(usernameDir) {
		log.Positive(fmt.Sprintf("Directory for GitHub user '%s' found!\nAdding new repository '%s'.", username, repoName))
	} else {
		log.Positive(fmt.Sprintf("Directory for GitHub user '%s' not found!\nCreating directory.", username))
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
