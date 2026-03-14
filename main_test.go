// Package main provides the core functionality for updating GitHub star counts in Markdown and AsciiDoc files.
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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
		{"Project4 (⭐1k) (extra notes)", "Project4 (extra notes)"},
	}

	for _, test := range tests {
		output := removeStarsInfo(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}

func TestGetAccessToken(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", "test_token")

	token, err := getAccessToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "test_token" {
		t.Errorf("Expected 'test_token', got '%s'", token)
	}
}

func TestGetAccessTokenMissing(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", "")

	_, err := getAccessToken()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestParseRepoName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantOwner string
		wantRepo  string
		wantErr   bool
	}{
		{
			name:      "Simple owner/repo",
			input:     "google/go-github",
			wantOwner: "google",
			wantRepo:  "go-github",
		},
		{
			name:      "With trailing path segments",
			input:     "owner/repo/tree/main",
			wantOwner: "owner",
			wantRepo:  "repo",
		},
		{
			name:    "Single segment",
			input:   "owner",
			wantErr: true,
		},
		{
			name:    "Empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "Missing repo",
			input:   "owner/",
			wantErr: true,
		},
		{
			name:    "Missing owner",
			input:   "/repo",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			owner, repo, err := parseRepoName(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if owner != tt.wantOwner || repo != tt.wantRepo {
				t.Errorf("expected (%s, %s), got (%s, %s)", tt.wantOwner, tt.wantRepo, owner, repo)
			}
		})
	}
}

func TestMarkdownFindRepos(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "Single link",
			content:  "- [Project](https://github.com/owner/repo)",
			expected: []string{"https://github.com/owner/repo"},
		},
		{
			name:     "Multiple links",
			content:  "[A](https://github.com/a/b) and [B](https://github.com/c/d)",
			expected: []string{"https://github.com/a/b", "https://github.com/c/d"},
		},
		{
			name:     "Link with existing stars",
			content:  "[Repo (⭐100)](https://github.com/owner/repo)",
			expected: []string{"https://github.com/owner/repo"},
		},
		{
			name:     "Non-GitHub link ignored",
			content:  "[Docs](https://docs.example.com/guide)",
			expected: []string{},
		},
		{
			name:     "No links",
			content:  "Just some text without links.",
			expected: []string{},
		},
	}

	updater := &MarkdownUpdater{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updater.FindRepos(tt.content)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("index %d: expected %q, got %q", i, tt.expected[i], got[i])
				}
			}
		})
	}
}

// Helper to simulate the update flow for tests.
func runUpdateFlow(t *testing.T, content string, client *github.Client, updater LinkUpdater) string {
	t.Helper()

	repos, err := updater.FindRepos(content)
	if err != nil {
		t.Fatalf("FindRepos failed: %v", err)
	}

	stars := make(map[string]int)
	ctx := context.Background()
	for _, repoURL := range repos {
		count, fetchErr := getStarsCount(ctx, client, repoURL)
		if fetchErr != nil {
			t.Fatalf("getStarsCount failed for %s: %v", repoURL, fetchErr)
		}
		stars[repoURL] = count
	}

	updated, err := updater.UpdateContent(content, stars)
	if err != nil {
		t.Fatalf("UpdateContent failed: %v", err)
	}
	return updated
}

func TestUpdateStarCounts(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/testowner/testrepo", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"stargazers_count": 42}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := github.NewClient(server.Client())
	baseURL, _ := url.Parse(server.URL + "/")
	client.BaseURL = baseURL

	md := "- [TestRepo](https://github.com/testowner/testrepo)"
	updated := runUpdateFlow(t, md, client, &MarkdownUpdater{})
	expected := "- [TestRepo (⭐42)](https://github.com/testowner/testrepo)"
	if updated != expected {
		t.Errorf("expected %q, got %q", expected, updated)
	}
}

func TestUpdateStarCountsMultiple(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/owner/repo1", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"stargazers_count": 1}`)
	})
	mux.HandleFunc("/repos/owner/repo2", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"stargazers_count": 2}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := github.NewClient(server.Client())
	baseURL, _ := url.Parse(server.URL + "/")
	client.BaseURL = baseURL

	md := "- [R1](https://github.com/owner/repo1)\n- [R2](https://github.com/owner/repo2)"
	updated := runUpdateFlow(t, md, client, &MarkdownUpdater{})
	expected := "- [R1 (⭐1)](https://github.com/owner/repo1)\n- [R2 (⭐2)](https://github.com/owner/repo2)"
	if updated != expected {
		t.Errorf("expected %q, got %q", expected, updated)
	}
}

func TestUpdateStarCountsExistingStars(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/owner/repo", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"stargazers_count": 10}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := github.NewClient(server.Client())
	baseURL, _ := url.Parse(server.URL + "/")
	client.BaseURL = baseURL

	md := "- [R (⭐5)](https://github.com/owner/repo)"
	updated := runUpdateFlow(t, md, client, &MarkdownUpdater{})
	expected := "- [R (⭐10)](https://github.com/owner/repo)"
	if updated != expected {
		t.Errorf("expected %q, got %q", expected, updated)
	}
}

func TestUpdateStarCountsAsciiDoc(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/testowner/adoc-repo", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"stargazers_count": 99}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := github.NewClient(server.Client())
	baseURL, _ := url.Parse(server.URL + "/")
	client.BaseURL = baseURL

	adoc := "link:https://github.com/testowner/adoc-repo[My Repo]"
	updated := runUpdateFlow(t, adoc, client, &ASCIIDocUpdater{})
	expected := "link:https://github.com/testowner/adoc-repo[My Repo (⭐99)]"
	if updated != expected {
		t.Errorf("expected %q, got %q", expected, updated)
	}
}
