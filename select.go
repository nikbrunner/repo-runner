package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func selectRepo(repoBasePath string) string {
	log := NewLogUtil()
	repos, err := getRepos(repoBasePath)
	if err != nil {
		log.Negative("Error getting repositories: %s", err)
		os.Exit(1)
	}

	return fzf(repos, "Select repository: ")
}

func getRepos(repoBasePath string) ([]string, error) {
	usernames, err := readDirs(repoBasePath)
	if err != nil {
		return nil, fmt.Errorf("reading username directory: %w", err)
	}

	var repos []string
	for _, username := range usernames {
		repoPath := filepath.Join(repoBasePath, username)
		reposInUser, err := readDirs(repoPath)
		if err != nil {
			return nil, fmt.Errorf("reading repo directory for '%s': %w", username, err)
		}

		for _, repo := range reposInUser {
			repoEntry := filepath.Join(username, repo)
			repos = append(repos, repoEntry)
		}
	}

	return repos, nil
}
