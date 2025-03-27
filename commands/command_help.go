package commands

import (
	"fmt"

	"github.com/levabu/pokedexcli/internal/pokeapi"
)

func commandHelp(cfg *pokeapi.Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range GetCommands() {
		fmt.Printf("%v: %v\n", c.Name, c.Description)
	}
	return nil
}
