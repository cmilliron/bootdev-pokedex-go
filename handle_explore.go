package main

import (
	"fmt"

	"github.com/cmilliron/bootdev-pokedex-go/internal/pokeapi"
)

func handleExploreCommand(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No args were passed.\n")
	}
	name := args[0]
	locationInfo, err := cfg.pokeapiClient.GetLocationDetails(name)
	if err != nil {
		return err
	}

	displayPokemon(locationInfo)


	return nil
}

func displayPokemon(locationInfo pokeapi.LocationAreaDetailResponse) {
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationInfo.Location.Name)
	for _, pokemon := range locationInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
}