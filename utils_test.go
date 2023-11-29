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
