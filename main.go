package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: repo [--open|--add|--remove|--status|--help]")
		return
	}

	switch os.Args[1] {
	case "--open":
		openRepo()
	case "--add":
		cloneRepo()
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

func cloneRepo() {
	fmt.Println("Cloning repository...")
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

func printNegative(message string) {
	fmt.Printf("%s%s%s\n", colorRed, message, colorReset)
}
