package main

func openRepo(config Config) {
	selectedRepo := selectRepo(config.ReposBasePath)
	sessionPath := createSessionPath(config.ReposBasePath, selectedRepo)
	sessionName := createSessionName(config.Separator, selectedRepo)
	attachToSession(sessionName, sessionPath)
}
