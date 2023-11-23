package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func selectRepo(basePath string) string {
	entries, err := os.ReadDir(basePath)
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

	if len(repos) == 0 {
		printNegative(fmt.Sprintf("No repositories found in %s", basePath), nil)
		os.Exit(1)
	}

	// Prepare the list of repositories as a newline-separated string
	repoList := strings.Join(repos, "\n")

	// Use fzf to select a repository
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(repoList)
	cmd.Stderr = os.Stderr // Connect fzf stderr to os.Stderr
	out, err := cmd.Output()
	if err != nil {
		printNegative("Error running fzf:", err)
		os.Exit(1)
	}

	selectedRepo := strings.TrimSpace(string(out))
	if selectedRepo == "" {
		printNegative("No repository selected", nil)
		os.Exit(1)
	}

	return selectedRepo
}
