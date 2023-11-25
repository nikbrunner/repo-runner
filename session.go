package main

import (
	"fmt"
	"os"
	"strings"
	"time"
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
	if err != nil {
		return false
	} else {
		return true
	}
}

func createSession(sessionName, path string) {
	if err := tmux("new-session", "-d", "-s", sessionName, "-c", path); err != nil {
		printNegative("Error creating new tmux session:", err)
		return
	}
	if err := tmux("rename-window", "-t", fmt.Sprintf("%s:1", sessionName), "code"); err != nil {
		printNegative("Error renaming tmux window to 'code':", err)
		return
	}
	if err := tmux("new-window", "-t", sessionName); err != nil {
		printNegative("Error creating new tmux window:", err)
		return
	}
	if err := tmux("rename-window", "-t", fmt.Sprintf("%s:2", sessionName), "run"); err != nil {
		printNegative("Error renaming tmux window to 'run':", err)
		return
	}
	if err := tmux("send-keys", "-t", fmt.Sprintf("%s:2", sessionName), "tmux_2x2_layout", "Enter"); err != nil {
		printNegative("Error setting up layout:", err)
		return
	}

	// Wait for tmux to create the layout and select the first window
	time.Sleep(2 * time.Second)
	if err := tmux("select-window", "-t", fmt.Sprintf("%s:1", sessionName)); err != nil {
		printNegative("Error selecting first tmux window:", err)
		return
	}

	printPositive("Session created")
}

func attachToSession(sessionName string, sessionPath string) {
	inTmux := os.Getenv("TMUX") != ""

	if sessionExists(sessionName) {
		if inTmux {
			if err := tmux("switch-client", "-t", sessionName); err != nil {
				printNegative("Error switching to tmux session:", err)
			}
		} else {
			if err := tmux("attach-session", "-t", sessionName); err != nil {
				printNegative("Error attaching to tmux session:", err)
			}
		}
	} else {
		printPositive("Creating session")
		createSession(sessionName, sessionPath)

		if inTmux {
			printPositive("Switching to session")
			if err := tmux("switch-client", "-t", sessionName); err != nil {
				printNegative("Error switching to tmux session:", err)
			}
		} else {
			printPositive("Attaching to session")
			if err := tmux("attach-session", "-t", sessionName); err != nil {
				printNegative("Error attaching to tmux session:", err)
			}
		}
	}
}
