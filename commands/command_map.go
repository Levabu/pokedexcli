package commands

import (
	"fmt"

	"github.com/levabu/pokedexcli/internal/pokeapi"
	"github.com/levabu/pokedexcli/internal/pokecache"
)

func commandMap(cfg *pokeapi.Config, args ...string) error {
	var url string
	if cfg.NextLocationsURL == nil {
		url = pokeapi.BaseURL + "/location-area"
	} else {
		url = *cfg.NextLocationsURL
	}
	return _baseMap(url, cfg)
}

func commandMapBack(cfg *pokeapi.Config, args ...string) error {
	if cfg.PreviousLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *cfg.PreviousLocationsURL
	return _baseMap(url, cfg)
}

func _baseMap(url string, cfg *pokeapi.Config) error {
	var locations pokeapi.RespLocations
	entry, isFound := cfg.PokeCache.Get(url)
	var err error
	if isFound {
		err = pokecache.ConvertCacheEntry(entry, &locations)
		if err != nil {
			return err
		}
	} else {
		locations, err = cfg.PokeClient.ListLocations(url, cfg)
		if err != nil {
			return err
		}
	}

	processLocations(&locations, cfg)
	return nil
}

func processLocations(locations *pokeapi.RespLocations, cfg *pokeapi.Config) {
	for _, res := range locations.Results {
		fmt.Println(res.Name)
	}

	cfg.PreviousLocationsURL = locations.Previous
	cfg.NextLocationsURL = locations.Next
}
