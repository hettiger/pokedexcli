package cmd

import (
	"fmt"
	"io"
)

func commandExit(out io.Writer) (Action, error) {
	_, err := fmt.Fprintln(out, "Closing the Pokedex... Goodbye!")

	return ActionExit, err
}
