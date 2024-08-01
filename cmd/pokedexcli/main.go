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
		if err := command.Callback(); err != nil {
			// Not sure if we should exit the program here. Will need to think about this more.
			log.Fatalf("Error running command '%s': %v", command.Name, err)
		}
		cliPrompt()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
}
