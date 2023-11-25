#!/bin/bash

# Default values
sessionName=""
sessionPath=""

# Process named arguments
while [[ "$#" -gt 0 ]]; do
	case $1 in
	--sessionName)
		sessionName="$2"
		shift
		;;
	--sessionPath)
		sessionPath="$2"
		shift
		;;
	*)
		echo "Unknown parameter: $1"
		exit 1
		;;
	esac
	shift
done

# Check if required arguments are provided
if [ -z "$sessionName" ] || [ -z "$sessionPath" ]; then
	echo "Missing arguments. Usage: $0 --sessionName <session-name> --sessionPath <session-path>"
	exit 1
fi

# Create a new tmux session
tmux new-session -d -s "$sessionName" -c "$sessionPath"

# Rename the first window to "code"
tmux rename-window -t "${sessionName}:1" "code"

# Create a second window
tmux new-window -t "$sessionName"

# Rename the second window to "run"
tmux rename-window -t "${sessionName}:2" "run"

# Now all commands should target window 2 of the new session

# Split the window into two equal vertical panes
tmux split-window -h -t "${sessionName}:2"

# Select the first pane and split it into two horizontal panes
tmux select-pane -t "${sessionName}:2.0"
tmux split-window -v -t "${sessionName}:2"

# Select the second pane (originally) and split it into two horizontal panes
tmux select-pane -t "${sessionName}:2.1"
tmux split-window -v -t "${sessionName}:2"

# Distribute all panes evenly
tmux select-layout -t "${sessionName}:2" tiled

# Send 'nvm use' and 'clear' to each pane
for pane in $(tmux list-panes -F '#P' -t "${sessionName}:2"); do
	tmux send-keys -t "${sessionName}:2.$pane" "nvm use" Enter clear Enter
done

# Select the first pane
tmux select-pane -t "${sessionName}:2.0"

# Select the first window
tmux select-window -t "${sessionName}:1"
