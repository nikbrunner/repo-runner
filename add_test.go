package main

import "testing"

func TestParseGitUrl(t *testing.T) {
	testCases := []struct {
		gitUrl       string
		wantUsername string
		wantRepoName string
		wantErr      bool
	}{
		{"git@github.com:user/repo.git", "user", "repo", false},
		{"https://github.com/user/repo.git", "user", "repo", false},
		{"invalid-url", "", "", true},
	}

	for _, tc := range testCases {
		gotUsername, gotRepoName, err := parseGitUrl(tc.gitUrl)
		if (err != nil) != tc.wantErr {
			t.Errorf("parseGitUrl(%s): expected error: %v, got: %v", tc.gitUrl, tc.wantErr, err)
		}
		if gotUsername != tc.wantUsername || gotRepoName != tc.wantRepoName {
			t.Errorf("parseGitUrl(%s): expected: %s/%s, got: %s/%s", tc.gitUrl, tc.wantUsername, tc.wantRepoName, gotUsername, gotRepoName)
		}
	}
}

func TestGetClonePath(t *testing.T) {
	testCases := []struct {
		gitUrl       string
		wantCloneDir string
		wantError    bool
	}{
		{"git@github.com:user/repo.git", "/home/user/repos/user/repo", false},
		{"https://github.com/user/repo.git", "/home/user/repos/user/repo", false},
		{"invalid-url", "", true},
	}

	config := Config{ReposBasePath: "/home/user/repos"}

	for _, tc := range testCases {
		gotCloneDir := getClonePath(tc.gitUrl, config)
		if gotCloneDir != tc.wantCloneDir {
			t.Errorf("getClonePath(%s): expected: %s, got: %s", tc.gitUrl, tc.wantCloneDir, gotCloneDir)
		}
	}
}
