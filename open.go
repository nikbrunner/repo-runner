package main

import (
	"fmt"
	"os"
)

func openRepo(config Config, repoName string) {
	log := NewLogUtil()
	repoBasePath := config.ReposBasePath
	var selectedRepo string

	if repoName == "" {
		selectedRepo = selectRepo(repoBasePath)
	} else {
		selectedRepo = repoName
	}

	sessionName := createSessionName(selectedRepo)
	sessionPath := createSessionPath(repoBasePath, selectedRepo)

	inTmux := os.Getenv("TMUX") != ""

	if sessionExists(sessionName) {
		if inTmux {
			log.Positive(fmt.Sprintf("Switching to session: %s", sessionName))
			if err := tmux("switch-client", "-t", sessionName); err != nil {
				log.Negative("Error switching to tmux session:", err)
			}
		} else {
			log.Positive(fmt.Sprintf("Attaching to session: %s", sessionName))
			if err := tmux("attach-session", "-t", sessionName); err != nil {
				log.Negative("Error attaching to tmux session:", err)
			}
		}
	} else {
		createSession(config, sessionName, sessionPath)

		if inTmux {
			log.Positive(fmt.Sprintf("Switching to session: %s", sessionName))
			if err := tmux("switch-client", "-t", sessionName); err != nil {
				log.Negative("Error switching to tmux session:", err)
			}
		} else {
			log.Positive(fmt.Sprintf("Attaching to session: %s", sessionName))
			if err := tmux("attach-session", "-t", sessionName); err != nil {
				log.Negative("Error attaching to tmux session:", err)
			}
		}
	}
}
