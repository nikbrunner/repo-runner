package main

import (
	"os"
	"os/exec"
	"strings"
)

func fzf(list []string, prompt string) string {
	// Define the color scheme (See tmux manual for --color)
	colorScheme := "fg:white,fg+:yellow,bg+:-1,gutter:-1,hl+:magenta,border:yellow,prompt:cyan,pointer:yellow,marker:cyan,spinner:green,header:blue,label:yellow,query:magenta"

	args := []string{
		"--reverse",
		"--no-separator",
		"--no-info",
		"--no-scrollbar",
		"--border=bold",
		"--border-label=┃   repo-runner ┃",
		"--border-label-pos=3",
		"--prompt", prompt,
		"--padding", "1,5",
		"--color", colorScheme,
	}

	cmd := exec.Command("fzf", args...)

	// Prepare the list of a newline-separated string
	preparedList := strings.Join(list, "\n")
	cmd.Stdin = strings.NewReader(preparedList)

	cmd.Stderr = os.Stderr // Connect fzf stderr to os.Stderr
	out, err := cmd.Output()
	// TODO: handle interrupt signal
	if err != nil {
		printNegative("Error running fzf:", err)
		os.Exit(1)
	}

	selectedRepo := strings.TrimSpace(string(out))
	if selectedRepo == "" {
		printNegative("No repository selected", nil)
		os.Exit(1)
	}

	return selectedRepo
}
