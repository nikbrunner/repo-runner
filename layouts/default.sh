#!/usr/bin/env bash

sessionName=$SESSION_NAME
sessionPath=$SESSION_PATH

# Create a new tmux session
tmux new-session -d -s "$sessionName" -c "$sessionPath"

# Rename the first window to "code"
tmux rename-window -t "${sessionName}:1" "code"

# Run "nvim ." in the first window
tmux send-keys -t "${sessionName}:1" "nvim ." Enter

# Create a second window in the same path
tmux new-window -t "$sessionName" -c "$sessionPath"

# Rename the second window to "run"
tmux rename-window -t "${sessionName}:2" "run"

# Split the window into two equal vertical panes
tmux split-window -h -t "${sessionName}:2" -c "$sessionPath"

# Select the first pane (now pane 1) and split it into two horizontal panes
tmux select-pane -t "${sessionName}:2.1"
tmux split-window -v -t "${sessionName}:2" -c "$sessionPath"

# Select the second pane (originally, now pane 2) and split it into two horizontal panes
tmux select-pane -t "${sessionName}:2.2"
tmux split-window -v -t "${sessionName}:2" -c "$sessionPath"

# Distribute all panes evenly
tmux select-layout -t "${sessionName}:2" tiled

# Send 'nvm use' and 'clear' to each pane
for pane in $(tmux list-panes -F '#P' -t "${sessionName}:2"); do
	tmux send-keys -t "${sessionName}:2.$pane" "nvm use" Enter clear Enter
done

# Select the first pane (now pane 1)
tmux select-pane -t "${sessionName}:2.1"

# Select the first window
tmux select-window -t "${sessionName}:1"
