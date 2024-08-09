package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const PokeAPIPokemonURL = "https://pokeapi.co/api/v2/pokemon/"

// Not sure if this is the best place to put the Pokedex; will need to think about this later
var Pokedex = map[string]Pokemon{}

// var catchProb = map[int]int{
// 	0:   90,
// 	100: 80,
// 	200: 70,
// 	300: 60,
// }

type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Moves struct {
	Move Move `json:"move"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Pokemon struct {
	Abilities      []Abilities `json:"abilities"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	Moves          []Moves     `json:"moves"`
	Name           string      `json:"name"`
	Species        Species     `json:"species"`
	Stats          []Stats     `json:"stats"`
	Types          []Types     `json:"types"`
	Weight         int         `json:"weight"`
}

// Will need to think about the logic a bit more and do a bit more experimenting.
// Will leave it like this for now since it does actually work.
func catchPokemon(baseExperience int) bool {
	chanceThresh := 0
	switch {
	case baseExperience >= 300:
		chanceThresh = 6
	case baseExperience >= 200:
		chanceThresh = 7
	case baseExperience >= 100:
		chanceThresh = 8
	default:
		chanceThresh = 9
	}

	rng := rand.Intn(10) + 1
	return rng <= chanceThresh
}

func CommandCatch(cmdInfo *CommandInfo, pokemonName string) error {
	var err error

	// Check that pokemonName isn't empty
	if pokemonName == "" {
		return fmt.Errorf("Please enter the name of a Pokemon")
	}

	// Check that pokemonName isn't a number
	_, err = strconv.ParseFloat(pokemonName, 64)
	if err == nil {
		return fmt.Errorf("Please ensure that the Pokemon name is not just a number")
	}

	var body []byte
	var found bool

	// Attempt to get data from cache before requesting from API
	// Not sure if we even need the cache since we're already planning to store the
	// data in the Pokedex. Will leave it for now.
	body, found = cmdInfo.Cache.Get(pokemonName)
	if !found {
		body, err = getAPIData(PokeAPIPokemonURL + pokemonName)
		if err != nil {
			return err
		}
	}

	// Unmarshall the data
	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return fmt.Errorf("Could not unmarshal json: %w", err)
	}

	// Add pokemon data to the cache
	if !found {
		cmdInfo.Cache.Add(pokemonName, body)
	}

	// Attempt to catch the pokemon
	caser := cases.Title(language.Und)
	fmt.Printf("Throwing a Pokeball at %s...\n", caser.String(pokemon.Name))
	if catchPokemon(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", caser.String(pokemon.Name))
		fmt.Println("You may now inspect it with the 'inspect' command.")
		Pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", caser.String(pokemon.Name))
	}
	return nil
}
