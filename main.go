package main

import (
	"time"

	"github.com/levabu/pokedexcli/internal/pokeapi"
	"github.com/levabu/pokedexcli/internal/pokecache"
)

func main() {
	cfg := pokeapi.Config{
		PokeClient: pokeapi.NewClient(5 * time.Second),
		PokeCache: *pokecache.NewCache(8 * time.Second),
	}

	startRepl(&cfg)
}
