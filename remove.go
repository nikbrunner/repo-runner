package main

import (
	"fmt"
	"os"
)

func removeRepo(config Config) {
	selectedRepo := selectRepo(config.ReposBasePath)
	sessionName := createSessionName(config.Separator, selectedRepo)
	sessionPath := createSessionPath(config.ReposBasePath, selectedRepo)

	// Ask the user to confirm
	printInfo(fmt.Sprintf("Are you sure you want to remove %s? [y/N]", sessionPath))
	var response string
	fmt.Scanln(&response)
	if response != "y" {
		printNegative("Aborting", nil)
		return
	}

	// Remove the dirdctory
	if err := os.RemoveAll(sessionPath); err != nil {
		printNegative("Error removing session", err)
		return
	} else {
		printPositive(fmt.Sprintf("Removed directory: %s", sessionPath))
	}

	// Check if a session exists for that repo
	if doesSessionExist(sessionName) {
		printPositive(fmt.Sprintf("Killing session: %s", sessionName))
		if err := tmux("kill-session", "-t", sessionName); err != nil {
			printNegative(fmt.Sprintf("Error killing session %s:", sessionName), err)
			return
		}
	} else {
		printInfo(fmt.Sprintf("No session found for %s. Skipping removal.", sessionName))
	}
}
