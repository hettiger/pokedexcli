package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hettiger/pokedexcli/internal/cmd"
)

func Run(in io.Reader, out io.Writer) error {
	commands := cmd.Commands()
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

		command, ok := commands[args[0]]
		if !ok {
			if _, err := fmt.Fprintln(out, "Unknown command"); err != nil {
				return err
			}

			continue
		}

		action, err := command.Callback(out)
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

		if action == cmd.ActionExit {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning for input failed: %w", err)
	}

	return nil
}
