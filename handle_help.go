package main

import "fmt"

func handleHelpCommand() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, command := range getCommands() {
		fmt.Printf("%s: %s\n", key, command.description)
	}
	return nil
}