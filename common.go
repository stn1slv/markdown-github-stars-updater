package main

import (
	"fmt"
	"regexp"
	"strings"
)

// LinkUpdater defines the interface for updating repo links in different file formats.
type LinkUpdater interface {
	// FindRepos finds all GitHub repository links in the given content.
	FindRepos(content string) ([]string, error)

	// UpdateContent updates the content by injecting star counts using the provided map.
	UpdateContent(content string, stars map[string]int) (string, error)
}

// removeStarsInfo removes the existing star count information from the input string.
func removeStarsInfo(input string) string {
	// Create a regular expression to find the "(⭐...)" pattern
	re := regexp.MustCompile(`\(⭐.*\)`)
	// Replace the matched substrings with an empty string
	result := re.ReplaceAllString(input, "")
	return strings.TrimSpace(result)
}

// formatStarCount formats the given star count for display in the markdown content.
func formatStarCount(stars int) string {
	if stars < 1000 {
		return fmt.Sprintf("%d", stars)
	} else if stars < 10000 {
		wholePart := stars / 1000
		decimalPart := (stars % 1000) / 100
		if decimalPart == 0 {
			return fmt.Sprintf("%dk", wholePart)
		}
		return fmt.Sprintf("%d.%dk", wholePart, decimalPart)
	} else {
		return fmt.Sprintf("%dk", stars/1000)
	}
}
