package cmd

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/hettiger/pokedexcli/internal/cfg"
)

func commandHelp(config *cfg.Config, out io.Writer) (cfg.Action, error) {
	var message strings.Builder

	message.WriteString("Welcome to the Pokedex!\nUsage:\n\n")

	commandNames := make([]string, 0, len(config.Commands))
	for name := range config.Commands {
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)

	for _, name := range commandNames {
		fmt.Fprintf(&message, "%s: %s\n", name, config.Commands[name].Description)
	}

	_, err := fmt.Fprint(out, message.String())

	return cfg.ActionContinue, err
}
