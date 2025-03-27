package commands

import (
	"fmt"

	"github.com/levabu/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *pokeapi.Config) error {
	return _baseMap(cfg.NextLocationsURL, cfg)
}

func commandMapBack(cfg *pokeapi.Config) error {
	if cfg.PreviousLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return _baseMap(cfg.PreviousLocationsURL, cfg)
}

func _baseMap(url *string, cfg *pokeapi.Config) error {
	locations, err := cfg.PokeClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.PreviousLocationsURL = locations.Previous
	cfg.NextLocationsURL = locations.Next

	for _, res := range locations.Results {
		fmt.Println(res.Name)
	}

	return nil
}
