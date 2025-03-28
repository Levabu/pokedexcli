package commands

import (
	"github.com/levabu/pokedexcli/internal/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 names of location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 names of location areas in the Pokemon world",
			Callback:    commandMapBack,
		},
		"explore": {
			Name: "explore",
			Description: "Displays the names of the pokemons found in a given location area",
			Callback: commandExplore,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
