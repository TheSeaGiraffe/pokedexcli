package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	PokeAPILocationsURL   = "https://pokeapi.co/api/v2/location-area/"
	PokeAPILocationsLimit = 20
)

type PokeAPIInfo struct {
	Next string
	Prev string
}

type PokeAPILocationResults struct {
	Name string `json:name`
	Url  string `json:url`
}

type PokeAPILocations struct {
	Count    int    `json:count`
	Next     string `json:next`
	Previous string `json:previous`
	Results  []PokeAPILocationResults
}

func unmarshallLocationDat(data []byte, locations *PokeAPILocations) error {
	err := json.Unmarshal(data, locations)
	if err != nil {
		return fmt.Errorf("Could not unmarshal json: %w", err)
	}

	return nil
}

func getLocationData(apiURL string) ([]byte, error) {
	res, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("Problem retrieving location data: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return nil, fmt.Errorf("Problem reading response body: %w", err)
	}

	return body, nil
}

func CommandMap(apiInfo *PokeAPIInfo) error {
	// Get data from API
	body, err := getLocationData(apiInfo.Next)
	if err != nil {
		return err
	}

	// Unmarshall the data
	locations := PokeAPILocations{}
	err = unmarshallLocationDat(body, &locations)
	if err != nil {
		return err
	}

	// Check to see if the results are empty
	if len(locations.Results) == 0 {
		return fmt.Errorf("No more results to display")
	}

	// Assign new Next and Previous URLs to apiInfo
	apiInfo.Next = locations.Next
	apiInfo.Prev = locations.Previous

	// Print the results
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}

// Honestly, probably want to just call Map with the url to the prev page.
func CommandMapb(apiInfo *PokeAPIInfo) error {
	// Check to see if we're already on the first page
	if apiInfo.Prev == "" {
		return fmt.Errorf("Already on the first page")
	}

	// Get data from the API
	body, err := getLocationData(apiInfo.Prev)
	if err != nil {
		return err
	}

	// Unmarshall the data
	locations := PokeAPILocations{}
	err = unmarshallLocationDat(body, &locations)
	if err != nil {
		return err
	}

	// Assign new Next and Previous URLs to apiInfo
	apiInfo.Next = locations.Next
	apiInfo.Prev = locations.Previous

	// Check to see if the results are empty
	// This should never happen but just in case
	if len(locations.Results) == 0 {
		return fmt.Errorf("Not allowed to move beyond the first page.")
	}

	// Print the results
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}
