// Package main provides the core functionality for updating GitHub star counts in Markdown and AsciiDoc files.
package main

import (
	"fmt"
	"regexp"
	"strings"
)

var markdownLinkRe = regexp.MustCompile(`\[([^\]]+)\]\((https://github\.com/[^/)]+/[^/)]+)\)`)

// MarkdownUpdater implements LinkUpdater for Markdown files.
type MarkdownUpdater struct{}

// FindRepos finds all GitHub repository links in the given content.
func (m *MarkdownUpdater) FindRepos(content string) ([]string, error) {
	matches := markdownLinkRe.FindAllStringSubmatch(content, -1)
	repos := make([]string, 0, len(matches))
	for _, match := range matches {
		repos = append(repos, match[2])
	}
	return repos, nil
}

// UpdateContent updates the content by injecting star counts using the provided map.
func (m *MarkdownUpdater) UpdateContent(content string, stars map[string]int) (string, error) {
	matches := markdownLinkRe.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		fullMatch := match[0]
		itemName := match[1]
		repoURL := match[2]

		starCount, ok := stars[repoURL]
		if !ok {
			continue
		}

		updatedLink := fmt.Sprintf("[%s (⭐%s)](%s)", removeStarsInfo(itemName), formatStarCount(starCount), repoURL)
		content = strings.Replace(content, fullMatch, updatedLink, 1)
	}
	return content, nil
}
