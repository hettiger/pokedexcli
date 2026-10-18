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
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of 20 location areas in the Pokemon world but in backward direction",
			Callback:    commandMapb,
		},
	}
}
