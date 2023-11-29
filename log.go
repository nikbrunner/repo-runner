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
	colorGray   = "\033[0;37m"
	colorReset  = "\033[0m"
)

const (
	symbolPositive = ""
	symbolInfo     = ""
	symbolNegative = "󰈸"
	symbolQuestion = ""
)

type Log struct{}

func NewLogUtil() *Log {
	return &Log{}
}

func (l *Log) formatMessage(symbol, color, message string) string {
	const separator = "  "
	return fmt.Sprintf("%s%s%s%s%s\n", color, symbol, separator, message, colorReset)
}

func (l *Log) Positive(message string) {
	fmt.Print(l.formatMessage(symbolPositive, colorGreen, message))
}

func (l *Log) Info(message string) {
	fmt.Print(l.formatMessage(symbolInfo, colorBlue, message))
}

func (l *Log) Neutral(message string) {
	fmt.Print(l.formatMessage(" ", colorGray, message))
}

func (l *Log) Negative(message string, err error) {
	fmt.Print(l.formatMessage(symbolNegative, colorRed, message))
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Log) Question(message string) {
	fmt.Print(l.formatMessage(symbolQuestion, colorOrange, message))
}

func (l *Log) Ask(s string) bool {
	var response string
	l.Question(fmt.Sprintf("%s (y/N): ", s))
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}
