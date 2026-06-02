package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageUrl *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

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
			return LocationAreasResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil || res.StatusCode != http.StatusOK {
			return LocationAreasResponse{}, err
		}
		defer res.Body.Close()

		dat, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreasResponse{}, err
		}
	
		data = dat
		c.cache.Add(url, data)
	}

	var locations LocationAreasResponse
	
	err := json.Unmarshal(data, &locations)
	if err != nil {
		// fmt.Errorf("Error decoding json: %v", err)
		return LocationAreasResponse{}, err
	}
	return locations, nil
}