// Package main provides the core functionality for updating GitHub star counts in Markdown and AsciiDoc files.
package main

import (
	"slices"
	"testing"
)

func TestAsciiDocFindRepos(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "Macro style link",
			content:  "Check out link:https://github.com/owner/repo[My Repo] here.",
			expected: []string{"https://github.com/owner/repo"},
		},
		{
			name:     "Inline style link",
			content:  "See https://github.com/owner/repo[My Repo] now.",
			expected: []string{"https://github.com/owner/repo"},
		},
		{
			name:     "Multiple links",
			content:  "link:https://github.com/a/b[AB] and https://github.com/c/d[CD]",
			expected: []string{"https://github.com/a/b", "https://github.com/c/d"},
		},
		{
			name:     "Link with existing stars",
			content:  "link:https://github.com/owner/repo[Repo (⭐100)]",
			expected: []string{"https://github.com/owner/repo"},
		},
		{
			name:     "Link with trailing path segments",
			content:  "link:https://github.com/owner/repo/tree/main[Browse Source]",
			expected: []string{"https://github.com/owner/repo/tree/main"},
		},
		{
			name:     "Link to issues page",
			content:  "https://github.com/owner/repo/issues[Issues]",
			expected: []string{"https://github.com/owner/repo/issues"},
		},
		{
			name:     "No links",
			content:  "Just some text without links.",
			expected: []string{},
		},
	}

	updater := &ASCIIDocUpdater{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updater.FindRepos(tt.content)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Equal(got, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestAsciiDocUpdateContent(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		stars    map[string]int
		expected string
	}{
		{
			name:     "Update macro link",
			content:  "link:https://github.com/owner/repo[Repo]",
			stars:    map[string]int{"https://github.com/owner/repo": 100},
			expected: "link:https://github.com/owner/repo[Repo (⭐100)]",
		},
		{
			name:     "Update inline link",
			content:  "https://github.com/owner/repo[Repo]",
			stars:    map[string]int{"https://github.com/owner/repo": 1500},
			expected: "https://github.com/owner/repo[Repo (⭐1.5k)]",
		},
		{
			name:     "Update existing stars",
			content:  "link:https://github.com/owner/repo[Repo (⭐50)]",
			stars:    map[string]int{"https://github.com/owner/repo": 100},
			expected: "link:https://github.com/owner/repo[Repo (⭐100)]",
		},
		{
			name:     "Mixed links",
			content:  "link:https://github.com/a/b[A] and https://github.com/c/d[B]",
			stars:    map[string]int{"https://github.com/a/b": 10, "https://github.com/c/d": 20},
			expected: "link:https://github.com/a/b[A (⭐10)] and https://github.com/c/d[B (⭐20)]",
		},
		{
			name:     "Update link with trailing path",
			content:  "link:https://github.com/owner/repo/tree/main[Source]",
			stars:    map[string]int{"https://github.com/owner/repo/tree/main": 42},
			expected: "link:https://github.com/owner/repo/tree/main[Source (⭐42)]",
		},
	}

	updater := &ASCIIDocUpdater{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updater.UpdateContent(tt.content, tt.stars)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
