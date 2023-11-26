package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration: ", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		getHelp()
		return
	}

	switch os.Args[1] {
	case "--open":
		openRepo(config, "")
	case "--add":
		cloneRepo(config, os.Args[2])
	case "--remove":
		removeRepo(config)
	case "--status":
		getStatus(config)
	case "--help":
		getHelp()
	default:
		fmt.Println("Invalid option. Usage: repo [--open|--add|--remove|--status|--help]")
	}
}
