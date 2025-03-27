package main

import (
	"time"

	"github.com/levabu/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := pokeapi.Config{
		PokeClient: pokeClient,
	}

	startRepl(&cfg)
}
