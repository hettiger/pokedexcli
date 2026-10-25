package tests

import (
	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/user"
)

func NewConfigFake(commands map[string]cfg.Command, pokeAPI *PokeAPIFake) *cfg.Config {
	return &cfg.Config{
		Commands: commands,
		PokeAPI:  pokeAPI,
		User:     user.NewUser(),
		Rand: func() float64 {
			return 33.33
		},
	}
}
