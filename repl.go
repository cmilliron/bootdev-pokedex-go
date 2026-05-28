package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		rawInput := scanner.Text()
		cleanedWords := NormalizeAndSplitInput(rawInput)

		if len(cleanedWords) == 0 {
			continue
		}

		command := cleanedWords[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}

func NormalizeAndSplitInput(text string) []string {
	lowercaseInput := strings.ToLower(text)

	words := strings.Fields(lowercaseInput)

	return words
}
