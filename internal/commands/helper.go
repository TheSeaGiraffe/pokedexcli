package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/TheSeaGiraffe/pokedexcli/internal/pokecache"
)

const (
	PokeAPILocationsURL   = "https://pokeapi.co/api/v2/location-area/"
	PokeAPILocationsLimit = 20
)

var ResponseFailedError = errors.New("Response Failed")

type CommandInfo struct {
	Next  string
	Prev  string
	Cache *pokecache.Cache
}

func NewCommandInfo() *CommandInfo {
	return &CommandInfo{
		Next:  PokeAPILocationsURL + fmt.Sprintf("?offset=0&limit=%d", PokeAPILocationsLimit),
		Cache: pokecache.NewCache(pokecache.CacheTTL),
	}
}

// Might want to use a stricter constraint here in the future but for now this will do.
func unmarshallAPIData[T any](data []byte, jsonStruct *T) error {
	err := json.Unmarshal(data, jsonStruct)
	if err != nil {
		return fmt.Errorf("Could not unmarshal json: %w", err)
	}

	return nil
}

func getAPIData(apiURL string) ([]byte, error) {
	res, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("Problem retrieving data: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		// return nil, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
		return nil, ResponseFailedError
	}
	if err != nil {
		return nil, fmt.Errorf("Problem reading response body: %w", err)
	}

	return body, nil
}
