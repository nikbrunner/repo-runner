package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func sanitizeSessionName(sessionName string) string {
	// https://unix.stackexchange.com/questions/560744/create-new-session-window-name-that-contain-dot
	return strings.NewReplacer(
		".", "_",
		":", "_",
	).Replace(sessionName)
}

func sanitizePath(path string) string {
	return strings.TrimSuffix(path, "/")
}

func createSessionName(repoPath string) string {
	santizedPath := sanitizePath(repoPath)
	repoName := strings.Split(santizedPath, "/")[len(strings.Split(santizedPath, "/"))-1]
	santizedSessionName := sanitizeSessionName(repoName)

	return santizedSessionName
}

func createSessionPath(basePath, repoName string) string {
	sanitizedPath := sanitizePath(basePath)
	return fmt.Sprintf("%s/%s", sanitizedPath, repoName)
}

func sessionExists(sessionName string) bool {
	err := tmux("has-session", "-t", sessionName)
	if err != nil {
		return false
	} else {
		printPositive(fmt.Sprintf("Session found: %s", sessionName))
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
