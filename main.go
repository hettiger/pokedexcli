package main

import (
	"log"
	"os"

	"github.com/hettiger/pokedexcli/internal/repl"
)

func main() {
	if err := repl.Run(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
