package main

import (
	"os"
	"os/exec"
	"strings"
)

func fzf(list []string, prompt string) string {
	args := []string{"--height", "50%", "--reverse", "--border", "--prompt", prompt}

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
