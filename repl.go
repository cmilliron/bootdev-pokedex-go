package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cmilliron/bootdev-pokedex-go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	pokeapiClient	*pokeapi.Client
	// pokeCache		*pokecache.Cache
	nextLocationURL	*string
	prevLocationURL	*string
}


func StartRepl(cfg *Config) {
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
			err := replCommand.callback(cfg)
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
		"map": {
			name: "map",
			description: "Displays the next 20 locations",
			callback: handleMapCommand,
		},
		"mapb": {
			name: "map",
			description: "Displays the previous 20 locations",
			callback: handleMapBCommand,
		},

	}
}