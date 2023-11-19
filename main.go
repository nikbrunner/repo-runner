package main

import (
	"fmt"
	"os"
)

type Config struct {
	RepoPath  string `json:"repoPath"`
	Separator string `json:"separator"`
}

func main() {
	if len(os.Args) < 2 {
		getHelp()
		return
	}

	switch os.Args[1] {
	case "--open":
		openRepo()
	case "--add":
		cloneRepo(os.Args[2])
	case "--remove":
		removeRepo()
	case "--status":
		getStatus()
	case "--help":
		getHelp()
	default:
		fmt.Println("Invalid option. Usage: repo [--open|--add|--remove|--status|--help]")
	}
}
