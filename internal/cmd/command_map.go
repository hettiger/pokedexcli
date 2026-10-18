package cmd

import (
	"fmt"
	"io"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/svc"
)

// commandMap displays the names of 20 location areas in the Pokemon world.
// Each subsequent call to map displays the next 20 locations.
func commandMap(config *cfg.Config, out io.Writer) (cfg.Action, error) {
	areas, err := config.PokeAPI.LocationAreas(svc.Forward)
	if err != nil {
		return cfg.ActionContinue, err
	}

	if len(areas.Results) == 0 {
		_, err := fmt.Fprintln(out, "you're on the last page")
		return cfg.ActionContinue, err
	}

	for _, area := range areas.Results {
		if _, err := fmt.Fprintln(out, area.Name); err != nil {
			return cfg.ActionContinue, err
		}
	}

	return cfg.ActionContinue, nil
}
