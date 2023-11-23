package main

import (
	"fmt"
	"strings"
)

func sanitizePath(path string) string {
	path = strings.TrimSuffix(path, "/")
	path = strings.TrimPrefix(path, "/")
	return path
}

func sanitizeSessionName(sessionName string) string {
	return strings.NewReplacer(
		".", "_",
		":", "_",
		"@", "_",
	).Replace(sessionName)
}

func createSessionName(separator string, repoPath string) string {
	sanatizedPath := sanitizePath(repoPath)
	parts := strings.Split(sanatizedPath, separator)
	sessionName := parts[len(parts)-1]

	return sanitizeSessionName(sessionName)
}

func createSessionPath(basePath, repoName string) string {
	return fmt.Sprintf("%s/%s", basePath, repoName)
}

func openRepo(config Config) {
	selectedRepo := selectRepo(config.ReposBasePath)
	sessionPath := createSessionPath(config.ReposBasePath, selectedRepo)
	sessionName := createSessionName(config.Separator, selectedRepo)
	attachToSession(sessionName, sessionPath)
}
