package pokeapi

type PokemonEncounters struct {
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}