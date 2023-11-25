package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type LastCommit struct {
	Author  string
	Message string
	Date    string
}

type RepoStatus struct {
	Name                  string
	Branch                string
	LastCommit            LastCommit
	HasUncommittedChanges bool
}

func getStatus(config Config) {
	statuses, err := getAllReposStatus(config.ReposBasePath)
	if err != nil {
		fmt.Printf("Error getting statuses: %v\n", err)
		os.Exit(1)
	}

	if len(statuses) > 0 {
		displaySummary(statuses)
		displayStatuses(statuses)

		if askForConfirmation("Open repositories with uncommitted changes?") {
			selectedRepo := fzf(getReposWithUncommittedChanges(statuses), "Select repository: ")
			openRepo(config, selectedRepo)
		} else if askForConfirmation("Open any other repository?") {
			openRepo(config, "")
		} else {
			printInfo("No repositories opened")
		}
	} else {
		printInfo("No repositories found")
	}
}

func getReposWithUncommittedChanges(statuses []RepoStatus) []string {
	var reposWithUncommmittedChanges []RepoStatus
	var repoNameList []string

	for _, status := range statuses {
		if status.HasUncommittedChanges {
			reposWithUncommmittedChanges = append(reposWithUncommmittedChanges, status)
		}
	}

	for _, repo := range reposWithUncommmittedChanges {
		repoNameList = append(repoNameList, repo.Name)
	}

	return repoNameList
}

func getAllReposStatus(reposBasePath string) ([]RepoStatus, error) {
	var statuses []RepoStatus

	err := filepath.Walk(reposBasePath, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() && isGitRepo(path) {
			status, err := getRepoStatus(path)
			if err != nil {
				return err
			}
			statuses = append(statuses, status)
		}
		return nil
	})

	return statuses, err
}

func isGitRepo(path string) bool {
	_, err := os.Stat(filepath.Join(path, ".git"))
	return !os.IsNotExist(err)
}

func getRepoStatus(repoPath string) (RepoStatus, error) {
	var status RepoStatus
	status.Name = filepath.Base(repoPath)

	// Change working directory to repo path
	oldPath, _ := os.Getwd()
	defer os.Chdir(oldPath)
	os.Chdir(repoPath)

	// Get branch and whether there are uncommitted changes
	branch, err := localGit("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return status, err
	}
	status.Branch = branch

	// Check if there are uncommitted changes
	uncommittedChanges, err := localGit("status", "--porcelain")
	if err != nil {
		return status, err
	}
	status.HasUncommittedChanges = uncommittedChanges != ""

	// Get Last Commit
	lastCommit, err := localGit("log", "-1", "--pretty=format:%an|%s|%cd")
	if err != nil {
		return status, err
	}

	// Parse Last Commit for Author, Message and Date
	lastCommitParts := strings.Split(lastCommit, "|")
	status.LastCommit.Author = lastCommitParts[0]
	status.LastCommit.Message = lastCommitParts[1]
	status.LastCommit.Date = lastCommitParts[2]

	return status, nil
}

// TODO: use `git()` from `git.go`
func localGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(output)), nil
}

func displaySummary(statuses []RepoStatus) {
	printInfo(fmt.Sprintf("Number of repositories: %d\n", len(statuses)))
}

func displayStatuses(statuses []RepoStatus) {
	for _, status := range statuses {
		if status.HasUncommittedChanges {
			fmt.Printf("%s%s [%s]%s (Uncommited Changes)", colorRed, status.Name, status.Branch, colorReset)
		} else {
			fmt.Printf("%s%s [%s]%s", colorGreen, status.Name, status.Branch, colorReset)
		}
		fmt.Println()

		fmt.Printf("    Author:  %s%s%s\n", colorOrange, status.LastCommit.Author, colorReset)
		fmt.Printf("    Message: %s%s%s\n", colorGreen, status.LastCommit.Message, colorReset)
		fmt.Printf("    Date:    %s%s%s\n", colorBlue, status.LastCommit.Date, colorReset)
		fmt.Println()
	}
}
