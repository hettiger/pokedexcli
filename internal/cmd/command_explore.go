package cmd

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/hettiger/pokedexcli/internal/cfg"
)

var ErrCommandExploreMissingLocation = errors.New("missing location")

func commandExplore(config *cfg.Config, out io.Writer, args ...string) (cfg.Action, error) {
	var message strings.Builder

	if len(args) < 1 {
		return cfg.ActionContinue, ErrCommandExploreMissingLocation
	}

	area, err := config.PokeAPI.LocationArea(args[0])
	if err != nil {
		return cfg.ActionContinue, err
	}

	fmt.Fprintf(&message, "Exploring %s...\nFound Pokémon:\n", args[0])

	if len(area.PokemonEncounters) == 0 {
		fmt.Fprintln(&message, "No Pokémon encounters in this area.")
	}

	for _, pokemonEncounter := range area.PokemonEncounters {
		fmt.Fprintf(&message, "- %s\n", pokemonEncounter.Pokemon.Name)
	}

	_, err = fmt.Fprint(out, message.String())

	return cfg.ActionContinue, err
}
