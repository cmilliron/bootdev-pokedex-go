package main

import (
	"fmt"
	"os"
)

func handleExitCommand(cfg *Config, args ...string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}