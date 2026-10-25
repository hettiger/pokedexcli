package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/hettiger/pokedexcli/internal/cfg"
)

func commandInspect(config *cfg.Config, out io.Writer, args ...string) (cfg.Action, error) {
	if len(args) < 1 {
		return cfg.ActionContinue, ErrMissingPokemonName
	}

	pokemonName := args[0]

	var message strings.Builder

	if pokemon, ok := config.User.Pokedex.Inspect(pokemonName); !ok {
		fmt.Fprintln(&message, "you have not caught that pokemon")
	} else {
		fmt.Fprintf(&message, "Name: %s\n", pokemon.Name)
		fmt.Fprintf(&message, "Height: %d\n", pokemon.Height)
		fmt.Fprintf(&message, "Weight: %d\n", pokemon.Weight)

		fmt.Fprintf(&message, "Stats:\n")
		for _, s := range pokemon.Stats {
			fmt.Fprintf(&message, "  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}

		fmt.Fprintf(&message, "Types:\n")
		for _, t := range pokemon.Types {
			fmt.Fprintf(&message, "  - %s\n", t.Type.Name)
		}
	}

	_, err := fmt.Fprint(out, message.String())

	return cfg.ActionContinue, err
}
