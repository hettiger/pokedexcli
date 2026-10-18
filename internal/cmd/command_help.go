package cmd

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func commandHelp(out io.Writer) (Action, error) {
	var message strings.Builder

	message.WriteString("Welcome to the Pokedex!\nUsage:\n\n")

	commands := Commands()

	commandNames := make([]string, 0, len(commands))
	for name := range commands {
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)

	for _, name := range commandNames {
		fmt.Fprintf(&message, "%s: %s\n", name, commands[name].Description)
	}

	_, err := fmt.Fprint(out, message.String())

	return ActionContinue, err
}
