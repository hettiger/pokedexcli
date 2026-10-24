package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hettiger/pokedexcli/internal/cfg"
)

func Run(config *cfg.Config, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	for {
		if _, err := fmt.Fprint(out, "Pokedex > "); err != nil {
			return err
		}

		if ok := scanner.Scan(); !ok {
			break
		}

		args := cleanInput(scanner.Text())

		if len(args) < 1 {
			continue
		}

		command, ok := config.Commands[args[0]]
		if !ok {
			if _, err := fmt.Fprintln(out, "Unknown command"); err != nil {
				return err
			}

			continue
		}

		action, err := command.Callback(config, out, args[1:]...)
		if err != nil {
			if _, writeErr := fmt.Fprintf(
				out,
				"Error executing command %q: %v\n",
				command.Name,
				err,
			); writeErr != nil {
				return writeErr
			}
		}

		if action == cfg.ActionExit {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning for input failed: %w", err)
	}

	return nil
}
