package main

import "testing"

func TestIsSshUrl(t *testing.T) {
	testCases := []struct {
		url      string
		wantBool bool
	}{
		{"git@github.com:user/repo.git", true},
		{"invalid-url", false},
	}

	for _, tc := range testCases {
		gotBool := isSshUrl(tc.url)
		if gotBool != tc.wantBool {
			t.Errorf("isSshUrl(%s): expected: %v, got: %v", tc.url, tc.wantBool, gotBool)
		}
	}
}

func TestIsHttpUrl(t *testing.T) {
	testCases := []struct {
		url      string
		wantBool bool
	}{
		{"https://github.com/user/repo.git", true},
		{"invalid-url", false},
	}

	for _, tc := range testCases {
		gotBool := isHttpUrl(tc.url)
		if gotBool != tc.wantBool {
			t.Errorf("isHttpUrl(%s): expected: %v, got: %v", tc.url, tc.wantBool, gotBool)
		}
	}
}

func TestIsValidGitUrl(t *testing.T) {
	testCases := []struct {
		url      string
		wantBool bool
	}{
		{"git@github.com:user/repo.git", true},
		{"https://github.com/user/repo.git", true},
		{".gitinvalid-url", false},
	}

	for _, tc := range testCases {
		gotBool := isValidGitUrl(tc.url)
		if gotBool != tc.wantBool {
			t.Errorf("isValidGitUrl(%s): expected: %v, got: %v", tc.url, tc.wantBool, gotBool)
		}
	}
}

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
