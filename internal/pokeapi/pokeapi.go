package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Config struct {
	PokeClient           Client
	PreviousLocationsURL *string
	NextLocationsURL     *string
}
