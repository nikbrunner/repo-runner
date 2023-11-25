package main

import (
	"os"
	"os/exec"
)

func tmux(args ...string) error {
	// exec.Command creates a new process to run the tmux command with the provided arguments.
	cmd := exec.Command("tmux", args...)

	// cmd.Stdin = os.Stdin sets the standard input of the new process to be the same as the
	// standard input of the Go program. This allows tmux to read input from the same terminal
	// or console that started the Go program. It's crucial for commands that require
	// user interaction or are expecting input from the terminal.
	cmd.Stdin = os.Stdin

	// cmd.Stdout and cmd.Stderr are set to the standard output and standard error of the Go program.
	// This means that any output or errors from the tmux command will be shown in the terminal where the Go program is running. It ensures that you can see what tmux is doing or any errors it encounters.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// cmd.Run() starts the tmux command and waits for it to finish.
	return cmd.Run()
}
