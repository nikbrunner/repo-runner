package main

import (
	"fmt"
	"os"
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
	fmt.Println("Getting status for all repositories...")
	statuses, err := getAllReposStatus(config.ReposBasePath)
	if err != nil {
		fmt.Printf("Error getting statuses: %v\n", err)
		os.Exit(1)
	}

	if len(statuses) > 0 {
		displaySummary(statuses)

		if len(getReposWithUncommittedChanges(statuses)) > 0 {
			if askForConfirmation("Open repositories with uncommitted changes?") {
				selectedRepo := selectRepo(config.ReposBasePath)
				openRepo(config, selectedRepo)
			}
		} else if askForConfirmation("Open any repository?") {
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

	progressCounter := 0

	err := filepath.Walk(reposBasePath, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() && isGitRepo(path) {
			status, err := getRepoStatus(path)
			if err != nil {
				fmt.Printf("Error getting status for %s: %v\n", path, err)
				return nil // Continue processing other repositories
			}
			displayStatuses(statuses)
			progressCounter++
			updateProgress(progressCounter)
			statuses = append(statuses, status)
		}
		return nil
	})

	return statuses, err
}

func updateProgress(progressCounter int) {
	fmt.Printf("Processed repositories: %d\r", progressCounter)
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

	// Check if the repository has any commits in its history
	if _, err := statusGitCmd(repoPath, "rev-list", "-n", "1", "--all"); err != nil {
		// No commits in the repository's history
		return RepoStatus{
			Name:                  filepath.Base(repoPath),
			Branch:                "N/A",
			LastCommit:            LastCommit{Author: "N/A", Message: "Repository empty or not initialized", Date: "N/A"},
			HasUncommittedChanges: false,
		}, nil
	}

	// Attempt to get the current branch name
	branch, branchErr := statusGitCmd(repoPath, "rev-parse", "--abbrev-ref", "HEAD")
	if branchErr != nil {
		// Handle the case where HEAD is ambiguous (no commits on the current branch)
		branch = "N/A"
	}
	status.Branch = branch

	// Check if there are uncommitted changes
	uncommittedChanges, err := statusGitCmd(repoPath, "status", "--porcelain")
	if err != nil {
		uncommittedChanges = ""
	}
	status.HasUncommittedChanges = uncommittedChanges != ""

	// Get Last Commit, handling current branch with no commits
	lastCommit, err := statusGitCmd(repoPath, "log", "-1", `--pretty=format:%an<%s<%cd`)
	if err != nil {
		// Current branch has no commits yet or error in getting last commit
		status.LastCommit = LastCommit{Author: "N/A", Message: "No commits on this branch", Date: "N/A"}
	} else {
		// Parse Last Commit for Author, Message and Date
		lastCommitParts := strings.Split(lastCommit, "<")
		if len(lastCommitParts) >= 3 {
			status.LastCommit.Author = lastCommitParts[0]
			status.LastCommit.Message = lastCommitParts[1]
			status.LastCommit.Date = lastCommitParts[2]
		}
	}

	return status, nil
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
