package cfg

import (
	"math/rand/v2"

	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/svc"
	"github.com/hettiger/pokedexcli/internal/user"
)

type Config struct {
	// Commands provides a cached registry of supported commands
	Commands map[string]Command

	// PokeAPI service
	PokeAPI PokeAPI

	// User provides the Pokédex store to keep track of caught Pokémon
	User user.User

	// Rand returns a pseudo-random number in the half-open interval [0.0,100.0).
	Rand func() float64
}

func NewConfig(commands map[string]Command, pokeAPI PokeAPI) *Config {
	return &Config{
		Commands: commands,
		PokeAPI:  pokeAPI,
		User:     user.NewUser(),
		Rand: func() float64 {
			return rand.Float64() * 100
		},
	}
}

type PokeAPI interface {
	LocationAreas(direction svc.Direction) (dto.APIResourceList[dto.LocationArea], error)
	LocationArea(name string) (dto.LocationArea, error)
	Pokemon(name string) (dto.Pokemon, error)
}
