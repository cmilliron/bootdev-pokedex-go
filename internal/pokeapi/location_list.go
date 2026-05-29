package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Client) GetLocations(pageUrl *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d", res.StatusCode)
		return LocationAreasResponse{}, err
	}

	var locations LocationAreasResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		fmt.Errorf("Error decoding json: %v", err)
		return LocationAreasResponse{}, err
	}
	return locations, nil
}