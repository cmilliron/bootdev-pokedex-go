package main

import (
	"fmt"

	// "io"
	// "log"

	"github.com/cmilliron/bootdev-pokedex-go/internal/pokeapi"
)



func handleMapCommand(cfg *Config) error {
	// fmt.Printf("Next: %s\nPrev: %v\n", cfg.Next, *cfg.Prev)

	locations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous
	displayLocation(locations)

	return nil
}
func handleMapBCommand(cfg *Config) error {
	locations, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous
	displayLocation(locations)

	return nil
}

func displayLocation(locations pokeapi.LocationAreasResponse) {
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}