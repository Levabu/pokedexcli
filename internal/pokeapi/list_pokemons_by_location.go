package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) ListPokemonsByLocationArea(url string, cfg *Config) (PokemonEncounters, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncounters{}, fmt.Errorf("networking problem: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return PokemonEncounters{}, fmt.Errorf("error fulfilling the request: %v", err)
	}
	defer res.Body.Close()

	code := res.StatusCode
	if code > 299 {
		return PokemonEncounters{}, fmt.Errorf("non-ok status code: %v", res.Status)
	}

	var data PokemonEncounters
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonEncounters{}, fmt.Errorf("error reading json data: %v", res.Status)
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return PokemonEncounters{}, fmt.Errorf("error decoding data: %v", res.Status)
	}
	cfg.PokeCache.Add(url, bytes)

	return data, nil
}