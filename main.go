package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func setExecutePermissions(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return os.Chmod(path, 0755) // Sets the file to readable and executable by everyone
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to set execute permissions: %v", err)
	}
}

func main() {
	setExecutePermissions("./layouts")

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
