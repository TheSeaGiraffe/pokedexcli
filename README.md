# Pokedex CLI

## Usage

The beginnings of a Pokemon CLI game. Players are currently only capable of very basic
navigation and catching Pokemon. Below is the list of available commands:

| command   | description                                                                                                         |
| --------- | ------------------------------------------------------------------------------------------------------------------- |
| `help`    | Displays a help message containing all of the available commands and a brief description                            |
| `exit`    | Exit the program                                                                                                    |
| `map`     | Displays 20 locations. Subsequent calls to `map` desplay the next 20 locations.                                     |
| `mapb`    | Same as `map` but in reverse.                                                                                       |
| `explore` | Displays more detailed information about the specified region                                                       |
| `catch`   | Attempt to catch the specified Pokemon                                                                              |
| `inspect` | Displays more detailed information about the specified Pokemon. Assumes that the Pokemon is already in your Pokedex |
| `pokedex` | Displays all of the Pokemon currently in your Pokedex                                                               |

## Build instructions

Make sure that you have Go 1.22 installed on your system and run `make build`. Afterwards,
you should see a `bin` directory in your current working directory containing the
`pokecli` binary which you can then run with `./bin/pokedexcli`.
