package repl_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hettiger/pokedexcli/internal/cmd"
	"github.com/hettiger/pokedexcli/internal/repl"
	"github.com/hettiger/pokedexcli/internal/tests"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Run("it can show help and exit", func(t *testing.T) {
		t.Parallel()

		config := tests.NewConfigFake(cmd.Commands(), tests.NewPokeAPIFake(t, nil))

		var out bytes.Buffer

		input := "help\nexit\ninvalid\n"

		help, err := os.ReadFile("../testdata/help.txt")
		require.NoError(t, err)

		want := fmt.Sprintf("Pokedex > %sPokedex > Closing the Pokedex... Goodbye!\n", help)

		err = repl.Run(config, strings.NewReader(input), &out)
		require.NoError(t, err)

		require.Equal(t, want, out.String())
	})

	t.Run("it warns on invalid commands", func(t *testing.T) {
		t.Parallel()

		config := tests.NewConfigFake(cmd.Commands(), tests.NewPokeAPIFake(t, nil))

		var out bytes.Buffer

		input := "invalid\n"
		want := "Pokedex > Unknown command\nPokedex > "

		err := repl.Run(config, strings.NewReader(input), &out)
		require.NoError(t, err)

		require.Equal(t, want, out.String())
	})
}
