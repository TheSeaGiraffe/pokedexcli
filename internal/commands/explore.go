package commands

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type PokeAPIPokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokeAPIEncounter struct {
	Pokemon PokeAPIPokemon `json:"pokemon"`
}

type PokeAPILocationInfo struct {
	PokemonEncounters []PokeAPIEncounter `json:"pokemon_encounters"`
}

func CommandExplore(cmdInfo *CommandInfo, locName string) error {
	var err error

	// Check that locName isn't empty
	if locName == "" {
		return fmt.Errorf("Please enter a region name")
	}

	// Check that locName isn't a number
	_, err = strconv.ParseFloat(locName, 64)
	if err == nil {
		return fmt.Errorf("Please ensure that the region name is not just a number")
	}

	var body []byte
	var found bool

	// Attempt to get data from cache before requesting from API
	body, found = cmdInfo.Cache.Get(locName)
	if !found {
		body, err = getAPIData(PokeAPILocationsURL + locName)
		if err != nil {
			return err
		}
	}

	// Unmarshall the data
	locationInfo := PokeAPILocationInfo{}
	err = json.Unmarshal(body, &locationInfo)
	if err != nil {
		return fmt.Errorf("Could not unmarshal json: %w", err)
	}

	// Add pokemon data to the cache
	if !found {
		cmdInfo.Cache.Add(locName, body)
	}

	// Print the results
	fmt.Printf("Exploring %s...\n", locName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
