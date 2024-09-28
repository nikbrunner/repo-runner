package main

import (
	"fmt"
	// "os"
	"os/exec"
	"strings"
)

// TODO: Should call an optional layout script if the user has defined one in the config file
// I will leave this commented out for now, just for reference
func createSession(_ Config, sessionName string, sessionPath string) {
	log := NewLogUtil()
	// cmd := exec.Command("/bin/bash", "-s")

	// Create session alone with no layout at first
	cmd := exec.Command("tmux", "new-session", "-d", "-s", sessionName, "-c", sessionPath)

	// TODO: If a layout is specified, use it

	// Set environment variables for the script
	// cmd.Env = append(os.Environ(),
	// 	fmt.Sprintf("SESSION_NAME=%s", sessionName),
	// 	fmt.Sprintf("SESSION_PATH=%s", sessionPath),
	// )

	// if config.Layout == LayoutDefault {
	// 	cmd.Stdin = strings.NewReader(defaultLayoutScript)
	// 	cmd.Stderr = os.Stderr
	// } else {
	// 	log.Negative(fmt.Sprintf("Invalid layout: %s", config.Layout), nil)
	// 	return
	// }

	if err := cmd.Run(); err != nil {
		log.Negative("Error executing layout script:", err)
		return
	}

	log.Positive(fmt.Sprintf("Created session: %s", sessionName))
}

func createSessionName(repoPath string) string {
	santizedPath := sanitizePath(repoPath)

	// From the repo path take the username + the repository name
	userName := strings.Split(santizedPath, "/")[len(strings.Split(santizedPath, "/"))-2]
	repoName := strings.Split(santizedPath, "/")[len(strings.Split(santizedPath, "/"))-1]

	// puth the username in front of the repo name
	sessionName := fmt.Sprintf("%s_%s", userName, repoName)
	santizedSessionName := sanitizeSessionName(sessionName)

	return santizedSessionName
}

func sanitizeSessionName(sessionName string) string {
	// https://unix.stackexchange.com/questions/560744/create-new-session-window-name-that-contain-dot
	return strings.NewReplacer(
		".", "_",
		":", "_",
	).Replace(sessionName)
}

func createSessionPath(basePath, repoName string) string {
	sanitizedPath := sanitizePath(basePath)
	return fmt.Sprintf("%s/%s", sanitizedPath, repoName)
}

func sessionExists(sessionName string) bool {
	log := NewLogUtil()
	err := tmux("has-session", "-t", sessionName)
	if err != nil {
		return false
	} else {
		log.Positive(fmt.Sprintf("Session found: %s", sessionName))
		return true
	}
}
