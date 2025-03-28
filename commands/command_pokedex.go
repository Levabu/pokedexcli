package commands

import (
	"fmt"
	"sort"

	"github.com/levabu/pokedexcli/internal/pokeapi"
)

func commandPokedex(cfg *pokeapi.Config, args ...string) error {
	var pokemons []pokeapi.Pokemon
	for _, p := range cfg.Pokedex {
		pokemons = append(pokemons, p)
	}
	if len(pokemons) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}

	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].Name < pokemons[j].Name
	})

	fmt.Println("Your pokedex:")
	for _, p := range pokemons {
		fmt.Printf("  - %s\n", p.Name)
	}
	return nil
}