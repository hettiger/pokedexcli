package cfg

import (
	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/svc"
)

type Config struct {
	// Commands provides a cached registry of supported commands
	Commands map[string]Command

	// PokeAPI service
	PokeAPI PokeAPI
}

type PokeAPI interface {
	LocationAreas(direction svc.Direction) (dto.APIResourceList[dto.LocationArea], error)
	LocationArea(name string) (dto.LocationArea, error)
}
