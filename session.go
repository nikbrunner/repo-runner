package main

import (
	"fmt"
	"os"
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
	cmd := exec.Command("/bin/bash", "-s")

	// Set environment variables for the script
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("SESSION_NAME=%s", sessionName),
		fmt.Sprintf("SESSION_PATH=%s", sessionPath),
	)

	if config.Layout == LayoutDefault {
		cmd.Stdin = strings.NewReader(defaultLayoutScript)
		cmd.Stderr = os.Stderr
	} else {
		printNegative(fmt.Sprintf("Invalid layout: %s", config.Layout), nil)
		return
	}

	if err := cmd.Run(); err != nil {
		printNegative("Error executing layout script:", err)
		return
	}

	printPositive(fmt.Sprintf("Created session: %s", sessionName))
}
