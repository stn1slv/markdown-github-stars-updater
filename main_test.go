package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/google/go-github/v68/github"
)

func TestFormatStarCount(t *testing.T) {
	tests := []struct {
		stars    int
		expected string
	}{
		{999, "999"},
		{1000, "1k"},
		{2501, "2.5k"},
		{4708, "4.7k"},
		{5038, "5k"},
		{6100, "6.1k"},
		{12000, "12k"},
		{78456, "78k"},
	}

	for _, test := range tests {
		result := formatStarCount(test.stars)
		if result != test.expected {
			t.Errorf("For %d stars, expected '%s' but got '%s'", test.stars, test.expected, result)
		}
	}
}

func TestRemoveStarsInfo(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Project1 (⭐1k)", "Project1"},
		{"Project2 (⭐1.5k)", "Project2"},
		{"Project3", "Project3"},
	}

	for _, test := range tests {
		output := removeStarsInfo(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}

func TestGetAccessToken(t *testing.T) {
	// Backup the existing environment variable and set a temporary one for the test
	oldToken := os.Getenv("GITHUB_TOKEN")
	os.Setenv("GITHUB_TOKEN", "test_token")

	token, err := getAccessToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "test_token" {
		t.Errorf("Expected 'test_token', got '%s'", token)
	}

	// Restore the original environment variable
	os.Setenv("GITHUB_TOKEN", oldToken)
}

func TestGetAccessTokenMissing(t *testing.T) {
	oldToken := os.Getenv("GITHUB_TOKEN")
	os.Unsetenv("GITHUB_TOKEN")

	_, err := getAccessToken()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	os.Setenv("GITHUB_TOKEN", oldToken)
}

func TestParseRepoName(t *testing.T) {
	owner, repo := parseRepoName("google/go-github")
	if owner != "google" || repo != "go-github" {
		t.Errorf("Expected google and go-github, got %s and %s", owner, repo)
	}
}

func TestUpdateStarCounts(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/testowner/testrepo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"stargazers_count": 42}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := github.NewClient(server.Client())
	baseURL, _ := url.Parse(server.URL + "/")
	client.BaseURL = baseURL

	md := "- [TestRepo](https://github.com/testowner/testrepo)"
	updated, err := updateStarCounts(md, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "- [TestRepo (⭐42)](https://github.com/testowner/testrepo)"
	if updated != expected {
		t.Errorf("expected %q, got %q", expected, updated)
	}
}
