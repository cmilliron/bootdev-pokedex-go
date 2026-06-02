package main

import (
	"time"

	"github.com/cmilliron/bootdev-pokedex-go/internal/pokeapi"
)

func main () {
	pokeClient := pokeapi.NewClient(5 * time.Second, 10 * time.Second)
	cfg := Config{
		pokeapiClient: pokeClient,
	}

	StartRepl(&cfg)
}