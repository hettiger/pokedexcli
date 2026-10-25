package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/hettiger/pokedexcli/internal/cfg"
)

func commandCatch(config *cfg.Config, out io.Writer, args ...string) (cfg.Action, error) {
	var message strings.Builder

	if len(args) < 1 {
		return cfg.ActionContinue, ErrMissingPokemonName
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
		config.User.Pokedex.Catch(pokemon)
		fmt.Fprintf(&message, "%s was caught!\n", pokemonName)
	} else {
		fmt.Fprintf(&message, "%s escaped!\n", pokemonName)
	}

	_, err = fmt.Fprint(out, message.String())

	return cfg.ActionContinue, err
}
