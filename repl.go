package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	commandRegistry := getCommands()

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		rawInput := scanner.Text()
		cleanedWords := NormalizeAndSplitInput(rawInput)

		if len(cleanedWords) == 0 {
			continue
		}

		commandName := cleanedWords[0]
		// fmt.Printf("Your command was: %s\n", commandName)
		replCommand, ok := commandRegistry[commandName]
		if ok {
			err := replCommand.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// Commands




func NormalizeAndSplitInput(text string) []string {
	lowercaseInput := strings.ToLower(text)

	words := strings.Fields(lowercaseInput)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: handleExitCommand,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: handleHelpCommand,
		},
	}
}