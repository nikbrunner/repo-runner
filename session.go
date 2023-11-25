package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func sanitizeSessionName(sessionName string) string {
	return strings.NewReplacer(
		".", "_",
		":", "_",
		"@", "_",
	).Replace(sessionName)
}

func sanitizePath(path string) string {
	return strings.TrimSuffix(path, "/")
}

func createSessionName(separator string, repoPath string) string {
	sanitizedPath := sanitizePath(repoPath)
	parts := strings.Split(sanitizedPath, separator)
	sessionName := parts[len(parts)-1]

	return sanitizeSessionName(sessionName)
}

func createSessionPath(basePath, repoName string) string {
	sanitizedPath := sanitizePath(basePath)
	return fmt.Sprintf("%s/%s", sanitizedPath, repoName)
}

func sessionExists(sessionName string) bool {
	err := tmux("has-session", "-t", sessionName)
	printPositive(fmt.Sprintf("Session found: %s", sessionName))
	if err != nil {
		return false
	} else {
		return true
	}
}

func createSession(config Config, sessionName string, sessionPath string) {
	args := []string{
		fmt.Sprintf("./layouts/%s.sh", config.Layout),
		"--sessionName", sessionName,
		"--sessionPath", sessionPath,
	}

	if err := exec.Command("/bin/bash", args...).Run(); err != nil {
		printNegative("Error executing layout script:", err)
		return
	}

	printPositive(fmt.Sprintf("Created session: %s", sessionName))
}
