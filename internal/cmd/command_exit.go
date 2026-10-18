package cmd

import (
	"fmt"
	"io"

	"github.com/hettiger/pokedexcli/internal/cfg"
)

func commandExit(config *cfg.Config, out io.Writer) (cfg.Action, error) {
	_, err := fmt.Fprintln(out, "Closing the Pokedex... Goodbye!")

	return cfg.ActionExit, err
}
