package main

import (
	"log"
	"os"
	"time"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/cmd"
	"github.com/hettiger/pokedexcli/internal/pokecache"
	"github.com/hettiger/pokedexcli/internal/repl"
	"github.com/hettiger/pokedexcli/internal/svc"
)

func main() {
	commands := cmd.Commands()
	cache := pokecache.NewCache(10 * time.Second)
	pokeAPI := svc.NewPokeAPI("https://pokeapi.co/api/v2", cache)
	config := cfg.NewConfig(commands, pokeAPI)

	if err := repl.Run(config, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
