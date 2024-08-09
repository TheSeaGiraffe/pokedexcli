package commands

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*CommandInfo, string) error
}

var CliCommandMap = map[string]CliCommand{
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Callback:    CommandHelp,
	},
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
	},
	"map": {
		Name:        "map",
		Description: "Display 20 locations. Subsequent calls will display the next 20 locations",
		Callback:    CommandMap,
	},
	"mapb": {
		Name:        "mapb",
		Description: "Same as map but in reverse",
		Callback:    CommandMapb,
	},
	"explore": {
		Name:        "explore",
		Description: "Explore the specified area",
		Callback:    CommandExplore,
	},
	"catch": {
		Name:        "catch",
		Description: "Attempt to catch the specified Pokemon",
		Callback:    CommandCatch,
	},
	"inspect": {
		Name:        "inspect",
		Description: "View information about the specified Pokemon. Assumes that the Pokemon has already been caught",
		Callback:    CommandInspect,
	},
	"pokedex": {
		Name:        "pokedex",
		Description: "View all of the Pokemon that you have caught",
		Callback:    CommandPokedex,
	},
}

func PrintUsageInfo() {
	for cmdName, cmd := range CliCommandMap {
		fmt.Printf("%-10s%s\n", cmdName, cmd.Description)
	}
	fmt.Println()
}

// Find a way to use the cliCommandMap to print the usage information
// I'll have to see how Lane does it later.
func CommandHelp(cmdInfo *CommandInfo, dummy string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	return nil
}

func CommandExit(cmdInfo *CommandInfo, dummy string) error {
	os.Exit(0)
	return nil
}

func CommandInspect(cmdInfo *CommandInfo, pokemonName string) error {
	// Check that pokemonName isn't empty
	if pokemonName == "" {
		return fmt.Errorf("Please enter the name of a Pokemon")
	}

	// Check that pokemonName isn't a number
	_, err := strconv.ParseFloat(pokemonName, 64)
	if err == nil {
		return fmt.Errorf("Please ensure that the Pokemon name is not just a number")
	}

	pokemon, ok := Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught that Pokemon")
	}
	caser := cases.Title(language.Und)
	fmt.Printf("Name: %s\n", caser.String(pokemon.Name))
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}
	return nil
}

func CommandPokedex(cmdInfo *CommandInfo, dummy string) error {
	if len(Pokedex) == 0 {
		return fmt.Errorf("You have not caught any Pokemon. Try catching at least one with the 'catch' command.")
	}
	caser := cases.Title(language.Und)
	fmt.Println("Your Pokedex:")
	for _, pokemon := range Pokedex {
		fmt.Printf("  - %s\n", caser.String(pokemon.Name))
	}
	return nil
}
