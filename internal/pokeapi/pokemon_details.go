package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	var data []byte

	// Check cashe for existing url
	cache, exists := c.cache.Get(url)
	if exists {
		fmt.Printf("=== From Cache === \n")
		data = cache
	} else {
		// fetch data if cach is empty
		fmt.Printf("\n=== From Api ===: \n")
		
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("request", err)
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil || res.StatusCode != http.StatusOK {
			fmt.Println("status", err)
			return Pokemon{}, err
		}
		defer res.Body.Close()

		dat, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("reading", err)
			return Pokemon{}, err
		}
	
		data = dat
		c.cache.Add(url, data)
	}

	var pokemon Pokemon
	
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		// fmt.Errorf("Error decoding json: %v", err)
		fmt.Println("unmarshalling", err)
		return Pokemon{}, err
	}
	return pokemon, nil
}