package main

import (
	"os"
)

func main() {
	log := NewLogUtil()
	config, err := loadConfig()
	if err != nil {
		log.Negative("Failed to load configuration: ", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		openRepo(config, "")
		return
	}

	switch os.Args[1] {
	case "--open":
		openRepo(config, "")
	case "--add":
		addRepo(config, os.Args[2])
	case "--remove":
		removeRepo(config)
	case "--status":
		getStatus(config)
	case "--help":
		getHelp()
	default:
		openRepo(config, "")
	}
}
