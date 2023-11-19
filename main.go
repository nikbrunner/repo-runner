package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	RepoPath  string `json:"repoPath"`
	Separator string `json:"separator"`
}

func loadConfig() Config {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)

	if err != nil {
		fmt.Println("Error decoding config file: ", err)
		os.Exit(1)
	}

	return config
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
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
		printHelp()
	default:
		fmt.Println("Invalid option. Usage: repo [--open|--add|--remove|--status|--help]")
	}
}

func expandPath(path string) string {
	if strings.Contains(path, "$HOME") {
		home, err := os.UserHomeDir()
		if err != nil {
			printNegative("Error getting home directory:", &err)
			return ""
		}
		return strings.Replace(path, "$HOME", home, 1)
	}
	return path
}

func parseGitUrl(gitUrl string) (string, string, error) {
	gitUrl = strings.TrimSuffix(gitUrl, ".git")

	var parts []string

	if strings.Contains(gitUrl, "git@") { // SSH format
		parts = strings.Split(gitUrl, ":")
		if len(parts) < 2 {
			return "", "", fmt.Errorf("invalid Git Url")
		}

		parts = strings.Split(parts[1], "/")
	} else { // Assume HTTP/HTTPS format
		parts = strings.Split(gitUrl, "/")
	}

	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid Git Url")
	}

	username := parts[len(parts)-2]
	repoName := parts[len(parts)-1]

	return username, repoName, nil
}

func cloneRepo(gitUrl string) {
	config := loadConfig()

	config.RepoPath = expandPath(config.RepoPath)
	if config.RepoPath == "" {
		return
	}

	separator := config.Separator
	if separator == "" {
		separator = "@"
	}

	// Parsing the Git URL to get username and repo name
	username, repoName, err := parseGitUrl(gitUrl)
	if err != nil {
		printNegative("Error parsing Git URL", &err)
	}

	fullPath := fmt.Sprintf("%s/%s%s%s", config.RepoPath, username, separator, repoName)

	printPositive("Cloning repostitory to " + fullPath)

	cmd := exec.Command("git", "clone", gitUrl, fullPath)
	err = cmd.Run()
	if err != nil {
		printNegative("Error cloning repository:", &err)
		return
	}

	printPositive("Repository cloned successfully")
}

func removeRepo() {
	fmt.Println("Removing repository...")
}

func openRepo() {
	fmt.Println("Opening repository...")
}

func getStatus() {
	fmt.Println("Getting status of repositories...")
}

func printHelp() {
	printInfo("Usage: repo [--open|--add|--remove|--status|--help]")
}

const (
	colorGreen  = "\033[0;32m"
	colorBlue   = "\033[0;34m"
	colorRed    = "\033[0;31m"
	colorOrange = "\033[0;33m"
	colorReset  = "\033[0m"
)

func printPositive(message string) {
	fmt.Printf("%s%s%s\n", colorGreen, message, colorReset)
}

func printInfo(message string) {
	fmt.Printf("%s%s%s\n", colorBlue, message, colorReset)
}

func printNegative(message string, err *error) {
	fmt.Printf("%s%s%s\n", colorRed, message, colorReset)
	if err != nil {
		fmt.Println(err)
	}
}
