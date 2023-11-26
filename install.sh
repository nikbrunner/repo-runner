#!/usr/bin/env bash

APPLICATION_NAME="rr"
INSTALL_PATH="/usr/local/bin"
LAYOUTS_DIR="layouts"

GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
ORANGE='\033[0;33m'
NC='\033[0m' # No Color

SYMBOL_POSITIVE=""
SYMBOL_NEGATIVE="󰈸"
SYMBOL_INFO=""
SYMBOL_QUESTION=""
SYMBOL_SEPARATOR="  "

print_positive() {
	printf "${GREEN}${SYMBOL_POSITIVE}${SYMBOL_SEPARATOR}%s${NC}\n" "$1"
}

print_info() {
	printf "${BLUE}${SYMBOL_INFO}${SYMBOL_SEPARATOR}%s${NC}\n" "$1"
}

print_negative() {
	printf "${RED}${SYMBOL_NEGATIVE}${SYMBOL_SEPARATOR}%s${NC}\n" "$1"
}

print_question() {
	printf "${ORANGE}${SYMBOL_QUESTION}${SYMBOL_SEPARATOR}%s${NC}\n" "$1"
}

# Check if required commands are available
for cmd in go sudo; do
	if ! command -v "$cmd" &>/dev/null; then
		print_negative "The required command '$cmd' is not installed."
		exit 1
	fi
done

# Build the Go application
print_info "Building the Go application..."
if ! go build -o "$APPLICATION_NAME"; then
	print_negative "Build failed."
	exit 1
fi

# Move the binary to the install path
print_info "Moving the binary to '$INSTALL_PATH'..."
if ! sudo mv "$APPLICATION_NAME" "$INSTALL_PATH"; then
	print_negative "Failed to move the binary to '$INSTALL_PATH'. Please check your permissions."
	exit 1
fi

# Set execute permissions on the layouts directory
print_info "Setting execute permissions on the layouts directory..."
if ! sudo chmod -R +x "$LAYOUTS_DIR"; then
	print_negative "Failed to set execute permissions on the '$LAYOUTS_DIR' directory."
	exit 1
fi

print_positive "Build and installation successful."
print_info "You can now run the application globally using the command '$APPLICATION_NAME'."
