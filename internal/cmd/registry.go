package cmd

import "github.com/hettiger/pokedexcli/internal/cfg"

func Commands() map[string]cfg.Command {
	return map[string]cfg.Command{
		"exit": {
			Name:        "exit",
			Signature:   "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Signature:   "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Signature:   "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Signature:   "mapb",
			Description: "Displays the names of 20 location areas in the Pokemon world but in backward direction",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Signature:   "explore <area_name>",
			Description: "Displays the names of all Pokémon found in a given area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Signature:   "catch <pokemon_name>",
			Description: "Attempt to catch a Pokémon by name",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Signature:   "inspect <pokemon_name>",
			Description: "Inspect a caught Pokémon by name",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Signature:   "pokedex",
			Description: "Prints a list of all the names of the Pokémon the user has caught.",
			Callback:    commandPokedex,
		},
	}
}
