package cmd

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/hettiger/pokedexcli/internal/cfg"
)

var ErrCommandCatchMissingPokemonName = errors.New("missing pokemon name")

func commandCatch(config *cfg.Config, out io.Writer, args ...string) (cfg.Action, error) {
	var message strings.Builder

	if len(args) < 1 {
		return cfg.ActionContinue, ErrCommandCatchMissingPokemonName
	}

	pokemonName := args[0]

	fmt.Fprintf(&message, "Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := config.PokeAPI.Pokemon(pokemonName)
	if err != nil {
		return cfg.ActionContinue, err
	}

	chance := pokemon.Chance()

	fmt.Fprintf(&message, "Chance: %.2f%%\n", chance)

	if config.Rand() < chance {
		config.User.Pokedex.Catch(pokemonName)
		fmt.Fprintf(&message, "%s was caught!\n", pokemonName)
	} else {
		fmt.Fprintf(&message, "%s escaped!\n", pokemonName)
	}

	fmt.Fprint(out, message.String())

	return cfg.ActionContinue, nil
}
