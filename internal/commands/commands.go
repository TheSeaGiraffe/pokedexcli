package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*PokeAPIInfo) error
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
}

func PrintUsageInfo() {
	for cmdName, cmd := range CliCommandMap {
		fmt.Printf("%s: %s\n", cmdName, cmd.Description)
	}
	fmt.Println()
}

// Find a way to use the cliCommandMap to print the usage information
// I'll have to see how Lane does it later.
func CommandHelp(apiInfo *PokeAPIInfo) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	// fmt.Println("help: Displays a help message")
	// fmt.Println("exit: Exit the Pokedex")
	fmt.Println()

	return nil
}

func CommandExit(apiInfo *PokeAPIInfo) error {
	os.Exit(0)
	return nil
}
