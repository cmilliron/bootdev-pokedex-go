package main

import (
	"fmt"

	"github.com/cmilliron/bootdev-pokedex-go/internal/pokeapi"
)

func handleInspectCommand(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: inspect <pokemon>")
	}

	pokemon := args[0]

	pokemonData, exist := cfg.pokedex[pokemon]
	if !exist {
		fmt.Printf("You have not caught a %s yet.\n", pokemon)
		return nil
	}

	displayStats(pokemonData)

	return nil
}

func displayStats(p pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range p.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)

	}

}