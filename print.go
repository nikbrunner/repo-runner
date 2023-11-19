package main

import "fmt"

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

func printNegative(message string, err error) {
	fmt.Printf("%s%s%s\n", colorRed, message, colorReset)
	if err != nil {
		fmt.Println(err)
	}
}
