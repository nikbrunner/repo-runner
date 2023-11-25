package main

import (
	"fmt"
	"os"
)

func openRepo(config Config, repoName string) {
	var selectedRepo string

	if repoName == "" {
		selectedRepo = selectRepo(config.ReposBasePath)
	} else {
		selectedRepo = repoName
	}

	sessionName := createSessionName(config.Separator, selectedRepo)
	sessionPath := createSessionPath(config.ReposBasePath, selectedRepo)

	inTmux := os.Getenv("TMUX") != ""

	if sessionExists(sessionName) {
		if inTmux {
			printPositive(fmt.Sprintf("Switching to session: %s", sessionName))
			if err := tmux("switch-client", "-t", sessionName); err != nil {
				printNegative("Error switching to tmux session:", err)
			}
		} else {
			printPositive(fmt.Sprintf("Attaching to session: %s", sessionName))
			if err := tmux("attach-session", "-t", sessionName); err != nil {
				printNegative("Error attaching to tmux session:", err)
			}
		}
	} else {
		createSession(config, sessionName, sessionPath)

		if inTmux {
			printPositive(fmt.Sprintf("Switching to session: %s", sessionName))
			if err := tmux("switch-client", "-t", sessionName); err != nil {
				printNegative("Error switching to tmux session:", err)
			}
		} else {
			printPositive(fmt.Sprintf("Attaching to session: %s", sessionName))
			if err := tmux("attach-session", "-t", sessionName); err != nil {
				printNegative("Error attaching to tmux session:", err)
			}
		}
	}
}
