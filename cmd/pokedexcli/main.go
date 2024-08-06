package main

import (
	"bufio"
	"errors"
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
	cmdInfo := commands.NewCommandInfo()
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

		if err := command.Callback(cmdInfo, commandArg); err != nil {
			if errors.Is(err, commands.ResponseFailedError) {
				var errMsg string
				switch {
				case strings.Contains(command.Name, "map"):
					errMsg = "Response failed when getting list of locations"
				case command.Name == "explore":
					errMsg = "Location does not exist"
				}
				fmt.Println(errMsg)
			} else {
				fmt.Println(err.Error())
			}
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
