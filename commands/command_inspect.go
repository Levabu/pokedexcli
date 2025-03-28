package commands

import (
	"fmt"

	"github.com/levabu/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *pokeapi.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: inspect <pokemon_name>")
	}
	pokemonName := args[0]

	pokemon, exists := cfg.Pokedex[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, _type := range pokemon.Types {
		fmt.Printf("  - %s\n", _type.Type.Name)
	}

	return nil
}
