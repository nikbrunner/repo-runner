package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func tmux(args ...string) {
	// exec.Command creates a new process to run the tmux command with the provided arguments.
	cmd := exec.Command("tmux", args...)

	// cmd.Stdin = os.Stdin sets the standard input of the new process to be the same as the
	// standard input of the Go program. This allows tmux to read input from the same terminal
	// or console that started the Go program. It's crucial for commands that require
	// user interaction or are expecting input from the terminal.
	cmd.Stdin = os.Stdin

	// cmd.Stdout and cmd.Stderr are set to the standard output and standard error of the Go program.
	// This means that any output or errors from the tmux command will be shown in the terminal
	// where the Go program is running. It ensures that you can see what tmux is doing or any errors it encounters.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// cmd.Run() starts the tmux command and waits for it to finish.
	// If there's an error in starting or running the command, it's handled here.
	if err := cmd.Run(); err != nil {
		printNegative("Error running tmux:", err)
		os.Exit(1)
	}
}

func doesSessionExist(sessionName string) bool {
	// TODO: use tmux() func
	cmd := exec.Command("tmux", "has-session", "-t", sessionName)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func createSession(sessionName, path string) {
	tmux("new-session", "-d", "-s", sessionName, "-c", path)
	tmux("rename-window", "-t", fmt.Sprintf("%s:1", sessionName), "code")
	tmux("new-window", "-t", sessionName)
	tmux("rename-window", "-t", fmt.Sprintf("%s:2", sessionName), "run")
	tmux("send-keys", "-t", fmt.Sprintf("%s:2", sessionName), "tmux_2x2_layout", "Enter")

	// Wait for tmux to create layout, and select the first window
	time.Sleep(2 * time.Second)
	tmux("select-window", "-t", fmt.Sprintf("%s:1", sessionName))

	printPositive("Session created")
}

func attachToSession(sessionName string, sessionPath string) {
	inTmux := os.Getenv("TMUX") != ""

	if doesSessionExist(sessionName) {
		if inTmux {
			tmux("switch-client", "-t", sessionName)
		} else {
			tmux("attach-session", "-t", sessionName)
		}
	} else {
		printPositive("Creating session")
		createSession(sessionName, sessionPath)

		if inTmux {
			printPositive("Switching to session")
			tmux("switch-client", "-t", sessionName)
		} else {
			printPositive("Attaching to session")
			tmux("attach-session", "-t", sessionName)
		}
	}
}
