package main

import (
	"log"
	"os"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/cmd"
	"github.com/hettiger/pokedexcli/internal/repl"
	"github.com/hettiger/pokedexcli/internal/svc"
)

func main() {
	config := &cfg.Config{
		Commands: cmd.Commands(),
		PokeAPI:  svc.NewPokeAPI("https://pokeapi.co/api/v2"),
	}

	if err := repl.Run(config, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
