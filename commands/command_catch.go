package commands

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/levabu/pokedexcli/internal/pokeapi"
	"github.com/levabu/pokedexcli/internal/pokecache"
)

func commandCatch(cfg *pokeapi.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: catch <pokemon_name>")
	}

	pokemonName := args[0]

	if _, inPokedex := cfg.Pokedex[pokemonName]; inPokedex {
		fmt.Printf("%s is already in your pokedex \n", pokemonName)
		return nil
	}

	url := pokeapi.BaseURL + "/pokemon/" + pokemonName
	
	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	var pokemon pokeapi.Pokemon
	entry, isFound := cfg.PokeCache.Get(url)
	var err error
	if isFound {
		err = pokecache.ConvertCacheEntry(entry, &pokemon)
		if err != nil {
			return err
		}
	} else {
		pokemon, err = cfg.PokeClient.GetPokemon(url, cfg)
		if err != nil {
			return err
		}
	}

	processPokemon(cfg, pokemon)

	return nil
}

func processPokemon(cfg *pokeapi.Config, pokemon pokeapi.Pokemon) {
	minChance := 0.1
	chance := math.Min(1, math.Max(minChance, 20 / float64(pokemon.BaseExperience)))
	if isCaught := chance > rand.Float64(); isCaught {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
}
