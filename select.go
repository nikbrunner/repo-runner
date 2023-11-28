package main

import (
	"fmt"
	"os"
	"path/filepath"
)

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

func selectRepo(repoBasePath string) string {
	repos, err := getRepos(repoBasePath)
	if err != nil {
		printNegative("Error getting repositories: %s", err)
		os.Exit(1)
	}

	return fzf(repos, "Select repository: ")
}
