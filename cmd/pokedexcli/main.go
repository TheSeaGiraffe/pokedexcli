package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

	var command commands.CliCommand
	var ok bool

	cliPrompt()
	for scanner.Scan() {
		commandName, commandArg := "", ""
		userInput := strings.Fields(scanner.Text())
		// Find a better way of doing this
		nUserInput := len(userInput)
		if nUserInput == 0 {
			cliPrompt()
			continue
		} else if nUserInput > 0 {
			commandName = userInput[0]
			if nUserInput > 1 {
				commandArg = userInput[1]
			}
		}
		command, ok = commands.CliCommandMap[commandName]
		if !ok {
			fmt.Printf("No such command\n")
			cliPrompt()
			continue
		}

		if err := command.Callback(cmdMapInfo, commandArg); err != nil {
			fmt.Println(err.Error())
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
