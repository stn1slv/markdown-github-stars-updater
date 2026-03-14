// Package main provides the core functionality for updating GitHub star counts in Markdown and AsciiDoc files.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v68/github"
	"golang.org/x/oauth2"
)

const githubURLPrefix = "https://github.com/"

var version = "dev"

func main() {
	outPath := flag.String("out", "", "output file path (defaults to input file)")
	dryRun := flag.Bool("dry-run", false, "print updated markdown to stdout")
	showVersion := flag.Bool("version", false, "show version info and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("markdown-github-stars-updater version %s\n", version)
		return
	}

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: markdown-github-stars-updater [flags] <path_to_file>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	filePath := flag.Arg(0)

	contentBytes, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading the file:", err)
		os.Exit(1)
	}
	content := string(contentBytes)

	token, err := getAccessToken()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	client := newGitHubClient(token)

	// Select Updater based on extension
	ext := strings.ToLower(filepath.Ext(filePath))
	var updater LinkUpdater

	switch ext {
	case ".md", ".markdown":
		updater = &MarkdownUpdater{}
	case ".adoc", ".asciidoc":
		updater = &ASCIIDocUpdater{}
	default:
		// Default to Markdown for backward compatibility if no extension matches, or error out?
		// Spec says "Tool automatically detects file type or applies appropriate parsing".
		// Let's assume Markdown if unknown for now, or maybe print a warning?
		// Given the project name is "markdown-github-stars-updater", defaulting to markdown seems safe,
		// but explicit support for asciidoc is added.
		// Let's print a warning and default to markdown for now, or just fail.
		// Failing is safer to avoid corrupting other files.
		fmt.Fprintf(os.Stderr, "Unsupported file extension: %s. Supported: .md, .markdown, .adoc, .asciidoc\n", ext)
		os.Exit(1)
	}

	// 1. Find Repos
	repos, err := updater.FindRepos(content)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error finding repositories:", err)
		os.Exit(1)
	}

	// 2. Fetch Stars
	stars := make(map[string]int)
	ctx := context.Background()
	for _, repoURL := range repos {
		if _, exists := stars[repoURL]; exists {
			continue
		}
		count, fetchErr := getStarsCount(ctx, client, repoURL)
		if fetchErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not fetch stars for %s: %v\n", repoURL, fetchErr)
			continue
		}
		stars[repoURL] = count
	}

	// 3. Update Content
	updatedContent, err := updater.UpdateContent(content, stars)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error updating content:", err)
		os.Exit(1)
	}

	if *dryRun {
		fmt.Println(updatedContent)
		return
	}

	output := filePath
	if *outPath != "" {
		output = *outPath
	}

	// G306: Expect WriteFile permissions to be 0600 or less (gosec)
	// We use 0644 because this is a documentation tool and the files are usually public/shared.
	err = os.WriteFile(output, []byte(updatedContent), 0o644) //nolint:gosec
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing updated file:", err)
		os.Exit(1)
	}

	fmt.Println("File updated successfully.")
}

// getStarsCount takes a GitHub repository URL and returns the current number of stars.
func getStarsCount(ctx context.Context, client *github.Client, repoURL string) (int, error) {
	if !strings.HasPrefix(repoURL, githubURLPrefix) {
		return 0, fmt.Errorf("invalid GitHub URL: %s", repoURL)
	}

	owner, repo, err := parseRepoName(repoURL[len(githubURLPrefix):])
	if err != nil {
		return 0, err
	}

	repository, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return 0, err
	}

	return repository.GetStargazersCount(), nil
}

// parseRepoName takes a path like "owner/repo" (possibly with trailing segments, query strings, or fragments)
// and returns the owner and repo parts.
func parseRepoName(repoPath string) (string, string, error) {
	// Strip query string (?...) and fragment (#...) before splitting
	if idx := strings.IndexAny(repoPath, "?#"); idx != -1 {
		repoPath = repoPath[:idx]
	}
	parts := strings.SplitN(repoPath, "/", 3) //nolint:mnd
	if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid repository path: %q", repoPath)
	}
	return parts[0], parts[1], nil
}

// getAccessToken retrieves the GitHub access token from the environment variable
func getAccessToken() (string, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return "", errors.New("missing GITHUB_TOKEN; set a GitHub personal access token")
	}
	return token, nil
}

func newGitHubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}
