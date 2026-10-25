package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/hettiger/pokedexcli/internal/cfg"
)

func commandPokedex(config *cfg.Config, out io.Writer, args ...string) (cfg.Action, error) {
	var message strings.Builder

	caught := config.User.Pokedex.Caught()

	if len(caught) == 0 {
		fmt.Fprintln(&message, "Your Pokedex is empty.")
	} else {
		fmt.Fprintln(&message, "Your Pokedex:")
		for _, p := range caught {
			fmt.Fprintf(&message, "  - %s\n", p.Name)
		}
	}

	_, err := fmt.Fprint(out, message.String())

	return cfg.ActionContinue, err
}
