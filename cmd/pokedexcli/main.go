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
	cmdMapInfo := commands.NewCommandMapInfo()
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

		// fmt.Printf("\nmap next before: '%s'\n", cmdMapInfo.Next)
		// fmt.Printf("map previous before: '%s'\n\n", cmdMapInfo.Prev)

		if err := command.Callback(cmdMapInfo); err != nil {
			// Won't exit the program. Will instead print error message
			// fmt.Printf("Error running command '%s': '%v'", command.Name, err)
			fmt.Println(err)
		}
		if command.Name == "help" {
			commands.PrintUsageInfo()
		}

		// fmt.Printf("\nmap next after: '%s'\n", cmdMapInfo.Next)
		// fmt.Printf("map previous after: '%s'\n\n", cmdMapInfo.Prev)

		cliPrompt()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
}
