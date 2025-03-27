package commands

import (
	"fmt"
	"os"

	"github.com/levabu/pokedexcli/internal/pokeapi"
)

func commandExit(cfg *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
