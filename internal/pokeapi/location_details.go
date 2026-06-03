package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetails(location string) (LocationAreaDetailResponse, error){
	url := baseURL + "/location-area/" + location
	// if pageUrl != nil {
	// 	url = *pageUrl
	// }

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
			return LocationAreaDetailResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil || res.StatusCode != http.StatusOK {
			return LocationAreaDetailResponse{}, err
		}
		defer res.Body.Close()

		dat, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaDetailResponse{}, err
		}
	
		data = dat
		c.cache.Add(url, data)
	}

	var locations LocationAreaDetailResponse
	
	err := json.Unmarshal(data, &locations)
	if err != nil {
		// fmt.Errorf("Error decoding json: %v", err)
		return LocationAreaDetailResponse{}, err
	}
	return locations, nil
}