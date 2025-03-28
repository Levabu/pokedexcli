package pokeapi

import (
	"github.com/levabu/pokedexcli/internal/pokecache"
)

const (
	BaseURL string = "https://pokeapi.co/api/v2"
)

type Config struct {
	PokeClient           Client
	PreviousLocationsURL *string
	NextLocationsURL     *string
	PokeCache            pokecache.Cache
}
