package main

import "testing"

func TestSanitizeSessionName(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
	}{
		{"normal-name", "normal-name"},
		{"name.with.dots", "name_with_dots"},
		{"name@with@at", "name@with@at"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sanitized := sanitizeSessionName(tc.name)
			if sanitized != tc.expected {
				t.Errorf("sanitizeSessionName(%s): expected: %v, got: %v", tc.name, tc.expected, sanitized)
			}
		})
	}
}

func TestSanitizePath(t *testing.T) {
	testCases := []struct {
		path     string
		expected string
	}{
		{"/home/user/repos/repo-owner/repo1", "/home/user/repos/repo-owner/repo1"},
		{"/home/user/repos/repo-owner/repo1/", "/home/user/repos/repo-owner/repo1"},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			sanitized := sanitizePath(tc.path)
			if sanitized != tc.expected {
				t.Errorf("sanitizePath(%s): expected: %v, got: %v", tc.path, tc.expected, sanitized)
			}
		})
	}
}

func TestCreateSesssionName(t *testing.T) {
	testCases := []struct {
		repoPath        string
		wantSessionName string
	}{
		{"/home/user/repos/repo-owner/repo1", "repo1"},
		{"/Users/joe/Documents/dev/repositories/terra-theme/terra-core.nvim", "terra-core_nvim"},
		{"/Users/joe/Documents/dev/repositories/terra-theme/terra-core.nvim/", "terra-core_nvim"},
	}

	for _, tc := range testCases {
		gotString := createSessionName(tc.repoPath)
		if gotString != tc.wantSessionName {
			t.Errorf("createSessionName(%s): expected: %v, got: %v", tc.repoPath, tc.wantSessionName, gotString)
		}
	}
}

func TestCreateSessionPath(t *testing.T) {
	testCases := []struct {
		basePath string
		repoName string
		expected string
	}{
		{"/home/user/repos", "repo1", "/home/user/repos/repo1"},
		{"/home/user/repos/", "repo1", "/home/user/repos/repo1"},
	}

	for _, tc := range testCases {
		t.Run(tc.basePath, func(t *testing.T) {
			sessionPath := createSessionPath(tc.basePath, tc.repoName)
			if sessionPath != tc.expected {
				t.Errorf("createSessionPath(%s, %s): expected: %v, got: %v", tc.basePath, tc.repoName, tc.expected, sessionPath)
			}
		})
	}
}
