package main

import (
	"strings"
)

func NormalizeAndSplitInput(text string) []string {
	lowercaseInput := strings.ToLower(text)

	words := strings.Fields(lowercaseInput)

	return words
}
