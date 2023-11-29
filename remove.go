package main

import (
	"fmt"
	"os"
)

func removeRepo(config Config) {
	log := NewLogUtil()
	repoBasePath := config.ReposBasePath
	selectedRepo := selectRepo(repoBasePath)
	sessionName := createSessionName(selectedRepo)
	sessionPath := createSessionPath(repoBasePath, selectedRepo)

	log.Ask(fmt.Sprintf("Are you sure you want to remove %s?", sessionPath))

	// Remove the directory
	if err := os.RemoveAll(sessionPath); err != nil {
		log.Negative("Error removing session", err)
		return
	} else {
		log.Positive(fmt.Sprintf("Removed directory: %s", sessionPath))
	}

	// Check if a session exists for that repo
	if sessionExists(sessionName) {
		log.Positive(fmt.Sprintf("Killing session: %s", sessionName))
		if err := tmux("kill-session", "-t", sessionName); err != nil {
			log.Negative(fmt.Sprintf("Error killing session %s:", sessionName), err)
			return
		}
	} else {
		log.Info(fmt.Sprintf("No session found for %s. Skipping removal.", sessionName))
	}
}
