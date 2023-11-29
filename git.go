package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Used in `add.go` - Will be merged/resolved in the future
func git(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Used in `status.go` - Will be merged/resolved in the future
func statusGitCmd(repoPath string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = repoPath // Set the working directory to the repo path

	// Create a buffer to capture standard error
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Execute the command
	output, err := cmd.Output()
	if err != nil {
		// Print the standard error output along with the error
		return "", fmt.Errorf("command error: %v, stderr: %s", err, stderr.String())
	}

	return strings.TrimSpace(string(output)), nil
}
