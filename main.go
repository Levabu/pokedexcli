package main

import (
	"time"

	"github.com/levabu/pokedexcli/internal/pokeapi"
	"github.com/levabu/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := pokeapi.Config{
		PokeClient: pokeClient,
		PokeCache: *pokecache.NewCache(5 * time.Second),
	}

	startRepl(&cfg)
}
