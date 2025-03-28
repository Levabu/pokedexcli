package commands

import (
	"fmt"

	"github.com/levabu/pokedexcli/internal/pokeapi"
	"github.com/levabu/pokedexcli/internal/pokecache"
)

func commandExplore(cfg *pokeapi.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: explore <area_name>")
	}

	locationArea := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + locationArea
	
	
	fmt.Printf("Exploring %s...\n", locationArea)

	var pokemons pokeapi.PokemonEncounters
	entry, isFound := cfg.PokeCache.Get(url)
	var err error
	if isFound {
		err = pokecache.ConvertCacheEntry(entry, &pokemons)
		if err != nil {
			return err
		}
	} else {
		pokemons, err = cfg.PokeClient.ListPokemonsByLocationArea(url, cfg)
		if err != nil {
			return err
		}
	}

	fmt.Println("Found Pokemons:")
	for _, pokemon := range pokemons.Pokemons {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
