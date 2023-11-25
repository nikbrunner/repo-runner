package main

import (
	"fmt"
	"strings"
)

const (
	colorGreen  = "\033[0;32m"
	colorBlue   = "\033[0;34m"
	colorRed    = "\033[0;31m"
	colorOrange = "\033[0;33m"
	colorReset  = "\033[0m"
)

const (
	symbolPositive = ""
	symbolInfo     = ""
	symbolNegative = "󰈸"
	symbolQuestion = ""
)

func symbol(symbol string) string {
	const separator = "  "
	return fmt.Sprintf("%s%s", symbol, separator)
}

func printPositive(message string) {
	fmt.Printf("%s%s%s%s\n", colorGreen, symbol(symbolPositive), message, colorReset)
}

func printInfo(message string) {
	fmt.Printf("%s%s%s%s\n", colorBlue, symbol(symbolInfo), message, colorReset)
}

func printNegative(message string, err error) {
	fmt.Printf("%s%s%s%s\n", colorRed, symbol(symbolNegative), message, colorReset)
	if err != nil {
		fmt.Println(err)
	}
}

func printQuestion(message string) {
	fmt.Printf("%s%s%s%s\n", colorOrange, symbol(symbolQuestion), message, colorReset)
}

func askForConfirmation(s string) bool {
	var response string
	printQuestion(fmt.Sprintf("%s (y/N): ", s))
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}
