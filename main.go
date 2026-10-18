package main

import (
	"log"
	"os"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/cmd"
	"github.com/hettiger/pokedexcli/internal/repl"
)

func main() {
	config := &cfg.Config{Commands: cmd.Commands()}

	if err := repl.Run(config, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
