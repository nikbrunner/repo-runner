package main

import (
	"os"
)

func getRepos(repoBasePath string) []string {
	entries, err := os.ReadDir(repoBasePath)
	if err != nil {
		printNegative("Error reading directory:", err)
		os.Exit(1)
	}

	var repos []string
	for _, entry := range entries {
		if entry.IsDir() {
			repos = append(repos, entry.Name())
		}
	}

	return repos
}

func selectRepo(repoBasePath string) string {
	return fzf(getRepos(repoBasePath), "Select repository: ")
}
