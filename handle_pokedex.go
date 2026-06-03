package main

import (
	"fmt"
)

func handlePokedexCommand(cfg *Config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caugh any Pokemon yet.")
		return nil
	}

	fmt.Println("Your Pokedex")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
