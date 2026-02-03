package main

import (
	"fmt"
	"regexp"
	"strings"
)

// AsciiDocUpdater implements LinkUpdater for AsciiDoc files.
type AsciiDocUpdater struct{}

func (a *AsciiDocUpdater) FindRepos(content string) ([]string, error) {
	// Pattern: (optional "link:") + (https://github.com/...) + [Text]
	// Group 1: URL
	// Group 2: Text (unused here but part of the structure)
	re := regexp.MustCompile(`(?:link:)?(https://github\.com/[^\[]+)\[([^\]]*)\]`)
	matches := re.FindAllStringSubmatch(content, -1)

	var repos []string
	for _, match := range matches {
		repos = append(repos, match[1])
	}
	return repos, nil
}

func (a *AsciiDocUpdater) UpdateContent(content string, stars map[string]int) (string, error) {
	re := regexp.MustCompile(`(?:link:)?(https://github\.com/[^\[]+)\[([^\]]*)\]`)
	matches := re.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		fullMatch := match[0]
		repoURL := match[1]
		text := match[2]

		starCount, ok := stars[repoURL]
		if !ok {
			continue
		}

		// Reconstruct the link with updated stars
		// We need to preserve the "link:" prefix if it was there
		prefix := ""
		if strings.HasPrefix(fullMatch, "link:") {
			prefix = "link:"
		}

		cleanText := removeStarsInfo(text)
		formattedStars := formatStarCount(starCount)
		newText := fmt.Sprintf("%s (⭐%s)", cleanText, formattedStars)

		updatedLink := fmt.Sprintf("%s%s[%s]", prefix, repoURL, newText)
		content = strings.Replace(content, fullMatch, updatedLink, 1)
	}
	return content, nil
}
