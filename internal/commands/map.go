package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TheSeaGiraffe/pokedexcli/internal/pokecache"
)

const (
	PokeAPILocationsURL   = "https://pokeapi.co/api/v2/location-area/"
	PokeAPILocationsLimit = 20
)

type CommandMapInfo struct {
	Next  string
	Prev  string
	Cache *pokecache.Cache
}

func NewCommandMapInfo() *CommandMapInfo {
	return &CommandMapInfo{
		Next:  PokeAPILocationsURL + fmt.Sprintf("?offset=0&limit=%d", PokeAPILocationsLimit),
		Cache: pokecache.NewCache(pokecache.CacheTTL),
	}
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

// Keeping this just in case I want to refactor the CommandMap and CommandMapb functions
// func commandMapBase(cmdMapInfo *CommandMapInfo, errorMsg string) error {
// 	// Get data from API
// 	body, err := getLocationData(cmdMapInfo.Next)
// 	if err != nil {
// 		return err
// 	}
//
// 	// Unmarshall the data
// 	locations := PokeAPILocations{}
// 	err = unmarshallLocationDat(body, &locations)
// 	if err != nil {
// 		return err
// 	}
//
// 	// Check to see if the results are empty
// 	if len(locations.Results) == 0 {
// 		return fmt.Errorf(errorMsg)
// 	}
//
// 	// Assign new Next and Previous URLs to apiInfo
// 	cmdMapInfo.Next = locations.Next
// 	cmdMapInfo.Prev = locations.Previous
//
// 	// Print the results
// 	for _, area := range locations.Results {
// 		fmt.Println(area.Name)
// 	}
//
// 	return nil
// }

// See if you can refactor ComandMap and ComamandMapb a bit more

func CommandMap(cmdMapInfo *CommandMapInfo) error {
	var body []byte
	var found bool
	var err error

	// Attempt to get data from cache before requesting from API
	body, found = cmdMapInfo.Cache.Get(cmdMapInfo.Next)
	if !found {
		body, err = getLocationData(cmdMapInfo.Next)
		if err != nil {
			return err
		}
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

	// Pass new values to cmdMapInfo
	// NOTE: don't move the block assigning the new locations to cmdMapInfo before the location data
	// is added to the cache as it will mess up the key-value pairing. You can figure out the logic if
	// you think about it for a bit.
	if !found {
		cmdMapInfo.Cache.Add(cmdMapInfo.Next, body) // This should add the url for the current page
	}
	cmdMapInfo.Next = locations.Next // This will update the url to that of the next page
	cmdMapInfo.Prev = locations.Previous

	// Print the results
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}

// Honestly, probably want to just call Map with the url to the prev page.
func CommandMapb(cmdMapInfo *CommandMapInfo) error {
	// Check to see if:
	// - The map command has been called
	// - We're already on the first page
	//
	// See if there's a better way of doing this
	if cmdMapInfo.Prev == "" {
		if cmdMapInfo.Next == PokeAPILocationsURL+fmt.Sprintf("?offset=0&limit=%d", PokeAPILocationsLimit) {
			return fmt.Errorf("Nothing to go back to. Try calling the 'map' command.")
		}
		return fmt.Errorf("Already on the first page")
	}

	var body []byte
	var found bool
	var err error

	// Attempt to get data from cache before requesting from API
	body, found = cmdMapInfo.Cache.Get(cmdMapInfo.Prev)
	if !found {
		body, err = getLocationData(cmdMapInfo.Prev)
		if err != nil {
			return err
		}
	}

	// Unmarshall the data
	locations := PokeAPILocations{}
	err = unmarshallLocationDat(body, &locations)
	if err != nil {
		return err
	}

	// Check to see if the results are empty
	if len(locations.Results) == 0 {
		return fmt.Errorf("Cannot move beyond the first page")
	}

	// Pass new values to cmdMapInfo
	// NOTE: don't move the block assigning the new locations to cmdMapInfo before the location data
	// is added to the cache as it will mess up the key-value pairing. You can figure out the logic if
	// you think about it for a bit.
	if !found {
		cmdMapInfo.Cache.Add(cmdMapInfo.Prev, body) // This should add the url for the current page
	}
	cmdMapInfo.Next = locations.Next
	cmdMapInfo.Prev = locations.Previous // We update the url to that of the previous page

	// Print the results
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}
