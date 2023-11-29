package main

import "fmt"

func getHelp() {
	log := NewLogUtil()

	log.Info("repo-runner")
	log.Neutral("A simple tool to manage and open your GitHub repositories.")
	fmt.Println()
	log.Neutral("Usage: rr [options]")
	fmt.Println()
	log.Neutral("--add <repo-url>    Add a repo")
	log.Neutral("--remove            Remove a repo")
	log.Neutral("--open              Open a repo")
	log.Neutral("--status            Show the status of all repos")
	log.Neutral("--help              Show this help message")
}
