package main

import (
	"fmt"
	"regexp"
	"strings"
)

type MarkdownUpdater struct{}

func (m *MarkdownUpdater) FindRepos(content string) ([]string, error) {
	re := regexp.MustCompile(`\[([^\]]+)\]\((https:\/\/github\.com\/[^\/)]+\/[^\/)]+)\)`)
	matches := re.FindAllStringSubmatch(content, -1)
	var repos []string
	for _, match := range matches {
		repos = append(repos, match[2])
	}
	return repos, nil
}

func (m *MarkdownUpdater) UpdateContent(content string, stars map[string]int) (string, error) {
	re := regexp.MustCompile(`\[([^\]]+)\]\((https:\/\/github\.com\/[^\/)]+\/[^\/)]+)\)`)
	matches := re.FindAllStringSubmatch(content, -1)

	// We replace one by one. To avoid issues with multiple same links, we should be careful.
	// The original code did `strings.Replace(..., match[0], ..., 1)` which iterates linearly.
	// This is safe if we iterate the same way.

	// However, modifying the string while iterating matches can be tricky if offsets change.
	// But `strings.Replace` searches for the substring.
	// Let's stick to the original logic for now which seemed to work for simple cases.

	for _, match := range matches {
		fullMatch := match[0]
		itemName := match[1]
		repoURL := match[2]

		starCount, ok := stars[repoURL]
		if !ok {
			// If we didn't fetch stars for this repo (e.g. error or skipped), keep original?
			// Or maybe we process all we found.
			continue
		}

		updatedLink := fmt.Sprintf("[%s (⭐%s)](%s)", removeStarsInfo(itemName), formatStarCount(starCount), repoURL)
		content = strings.Replace(content, fullMatch, updatedLink, 1)
	}
	return content, nil
}
