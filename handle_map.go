package main

import (
	"encoding/json"
	"fmt"

	// "io"
	// "log"
	"net/http"
)

type LocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func handleMapCommand(cfg *Config) error {
	// fmt.Printf("Next: %s\nPrev: %v\n", cfg.Next, *cfg.Prev)

	res, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	var locations LocationsResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return fmt.Errorf("Error decoding json: %v", err)
	}
	cfg.Next = locations.Next
	cfg.Prev = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
func handleMapBCommand(cfg *Config) error {
	// fmt.Printf("Next: %s\nPrev: %v\n", cfg.Next, *cfg.Prev)

	if cfg.Prev == nil {
		fmt.Println("No previous results")
		return nil
	}

	res, err := http.Get(*cfg.Prev)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	var locations LocationsResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return fmt.Errorf("Error decoding json: %v", err)
	}
	cfg.Next = locations.Next
	cfg.Prev = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}