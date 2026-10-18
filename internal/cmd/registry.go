package cmd

import "github.com/hettiger/pokedexcli/internal/cfg"

func Commands() map[string]cfg.Command {
	return map[string]cfg.Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
	}
}
