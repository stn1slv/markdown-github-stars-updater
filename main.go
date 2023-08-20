package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path_to_markdown_file>")
		return
	}
	// Read markdown content
	filePath := os.Args[1]

	markdownContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the markdown file:", err)
		return
	}

	// Parse markdown content
	updatedMarkdown, err := updateStarCounts(string(markdownContent))
	if err != nil {
		fmt.Println("Error updating star counts:", err)
		return
	}

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(updatedMarkdown), 0644)
	if err != nil {
		fmt.Println("Error writing updated markdown to file:", err)
		return
	}

	fmt.Println("Markdown file updated successfully.")
}

func updateStarCounts(markdownContent string) (string, error) {
	// Regular expression to find GitHub repository links
	re := regexp.MustCompile(`\[([^\]]+)\]\((https:\/\/github\.com\/[^\/)]+\/[^\/)]+)\)`)
	matches := re.FindAllStringSubmatch(markdownContent, -1)

	for _, match := range matches {
		itemName := match[1]
		repoURL := match[2]
		starCount, err := getStarsCount(repoURL)
		if err != nil {
			return "", err
		}

		// Update star count in the link title
		updatedLink := fmt.Sprintf("[%s (⭐%s)](%s)", removeStars(itemName), formatStarCount(starCount), repoURL)
		markdownContent = strings.Replace(markdownContent, match[0], updatedLink, 1)
	}

	return markdownContent, nil
}

func getStarsCount(repoURL string) (int, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: getAccessToken()},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	owner, repo := parseRepoName(repoURL[len("https://github.com/"):])
	repository, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return 0, err
	}

	return repository.GetStargazersCount(), nil
}

func parseRepoName(repoName string) (string, string) {
	parts := strings.Split(repoName, "/")
	if len(parts) != 2 {
		panic("Invalid repository name format")
	}
	return parts[0], parts[1]
}

func removeStars(input string) string {
	// Create a regular expression to find the "(⭐...)" pattern
	re := regexp.MustCompile(`\(⭐.*\)`)
	// Replace the matched substrings with an empty string
	result := re.ReplaceAllString(input, "")
	return strings.TrimSpace(result)
}

func formatStarCount(stars int) string {
	if stars < 1000 {
		return fmt.Sprintf("%d", stars)
	} else if stars < 10000 {
                // Check if the stars count is divisible by 1000
		rounded := float64(stars) / 1000
		if rounded == float64(int(rounded)) {
			return fmt.Sprintf("%.1fk", rounded)
		} else {
			return fmt.Sprintf("%.1fk", rounded)
		}
	} else {
		return fmt.Sprintf("%dk", stars/1000)
	}
}

func getAccessToken() string {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Error: GITHUB_TOKEN environment variable not set")
		os.Exit(1)
	}
	return token
}
