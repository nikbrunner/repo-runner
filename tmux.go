package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
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

func doesSessionExist(sessionName string) bool {
	err := tmux("has-session", "-t", sessionName)
	if err != nil {
		return false
	} else {
		return true
	}
}

func createSession(sessionName, path string) {
	if err := tmux("new-session", "-d", "-s", sessionName, "-c", path); err != nil {
		printNegative("Error creating new tmux session:", err)
		return
	}
	if err := tmux("rename-window", "-t", fmt.Sprintf("%s:1", sessionName), "code"); err != nil {
		printNegative("Error renaming tmux window to 'code':", err)
		return
	}
	if err := tmux("new-window", "-t", sessionName); err != nil {
		printNegative("Error creating new tmux window:", err)
		return
	}
	if err := tmux("rename-window", "-t", fmt.Sprintf("%s:2", sessionName), "run"); err != nil {
		printNegative("Error renaming tmux window to 'run':", err)
		return
	}
	if err := tmux("send-keys", "-t", fmt.Sprintf("%s:2", sessionName), "tmux_2x2_layout", "Enter"); err != nil {
		printNegative("Error setting up layout:", err)
		return
	}

	// Wait for tmux to create the layout and select the first window
	time.Sleep(2 * time.Second)
	if err := tmux("select-window", "-t", fmt.Sprintf("%s:1", sessionName)); err != nil {
		printNegative("Error selecting first tmux window:", err)
		return
	}

	printPositive("Session created")
}

func attachToSession(sessionName string, sessionPath string) {
	inTmux := os.Getenv("TMUX") != ""

	if doesSessionExist(sessionName) {
		if inTmux {
			if err := tmux("switch-client", "-t", sessionName); err != nil {
				printNegative("Error switching to tmux session:", err)
			}
		} else {
			if err := tmux("attach-session", "-t", sessionName); err != nil {
				printNegative("Error attaching to tmux session:", err)
			}
		}
	} else {
		printPositive("Creating session")
		createSession(sessionName, sessionPath)

		if inTmux {
			printPositive("Switching to session")
			if err := tmux("switch-client", "-t", sessionName); err != nil {
				printNegative("Error switching to tmux session:", err)
			}
		} else {
			printPositive("Attaching to session")
			if err := tmux("attach-session", "-t", sessionName); err != nil {
				printNegative("Error attaching to tmux session:", err)
			}
		}
	}
}
