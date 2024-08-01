package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/TheSeaGiraffe/pokedexcli/internal/commands"
)

const cliPromptName = "Pokedex"

func cliPrompt() {
	fmt.Printf("%s > ", cliPromptName)
}

// Wonder if there's a better way of handling all the CLI prompts
func main() {
	apiInfo := commands.PokeAPIInfo{
		Next: commands.PokeAPILocationsURL + fmt.Sprintf("?offset=0&limit=%d", commands.PokeAPILocationsLimit),
	}
	scanner := bufio.NewScanner(os.Stdin)
	cliPrompt()
	for scanner.Scan() {
		userCommand := scanner.Text()
		command, ok := commands.CliCommandMap[userCommand]
		if !ok {
			fmt.Println("No such command")
			cliPrompt()
			continue
		}
		if err := command.Callback(&apiInfo); err != nil {
			// Won't exit the program. Will instead print error message
			// fmt.Printf("Error running command '%s': '%v'", command.Name, err)
			fmt.Println(err)
		}
		if command.Name == "help" {
			commands.PrintUsageInfo()
		}

		cliPrompt()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
}
