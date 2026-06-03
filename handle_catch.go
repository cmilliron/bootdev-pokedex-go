package main

import (
	"fmt"
	"math/rand/v2"
)

func handleCatchCommand(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: catch <pokemon>")
	}

	pokemon := args[0]

	pokemonDetails, err := cfg.pokeapiClient.GetPokemonDetails(pokemon)
	if err != nil {
		return fmt.Errorf("There was an error fetching details for %s\n", pokemon)
	}

	baseExperience := pokemonDetails.BaseExperience
	chance := rand.N(200) + 50

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if baseExperience > chance {
		fmt.Printf("Oh oh, %s is too strong for you!\n", pokemon)
		
	} else {
		fmt.Printf("You caugh %s!\n", pokemon)
		cfg.pokedex[pokemon] = pokemonDetails
	} 
	
	return nil
}