package main

import (
	"testing"
)

func TestFormatStarCount(t *testing.T) {
	tests := []struct {
		stars    int
		expected string
	}{
		{999, "999"},
		{1000, "1k"},
		{2501, "2.5k"},
		{4108, "4.1k"},
		{6100, "6k"},
		{12000, "12k"},
	}

	for _, test := range tests {
		result := formatStarCount(test.stars)
		if result != test.expected {
			t.Errorf("For %d stars, expected '%s' but got '%s'", test.stars, test.expected, result)
		}
	}
}

func TestRemoveStar(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Project1 (⭐1k)", "Project1"},
		{"Project2 (⭐1.5k)", "Project2"},
		{"Project3", "Project3"},
	}

	for _, test := range tests {
		output := removeStars(test.input)
		if output != test.expected {
			t.Errorf("Input: %s\nExpected: %s\nGot: %s", test.input, test.expected, output)
		}
	}
}
