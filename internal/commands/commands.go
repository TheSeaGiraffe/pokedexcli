package commands

import (
	"fmt"
	"os"
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
}

func PrintUsageInfo() {
	for cmdName, cmd := range CliCommandMap {
		// Need to add padding here.
		// fmt.Printf("%s:%s\n", cmdName, cmd.Description)
		fmt.Printf("%-10s%s\n", cmdName, cmd.Description)
		// fmt.Printf("%-10s%s\n", cmdName+":", cmd.Description)
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
