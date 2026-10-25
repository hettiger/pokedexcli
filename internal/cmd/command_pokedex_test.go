package cmd

import (
	"bytes"
	"testing"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/tests"
	"github.com/stretchr/testify/require"
)

func Test_commandPokedex(t *testing.T) {
	config := tests.NewConfigFake(Commands(), tests.NewPokeAPIFake(t, nil))

	t.Run("it warns if pokedex is still empty", func(t *testing.T) {
		t.Parallel()

		var output bytes.Buffer

		got, gotErr := commandPokedex(config, &output)

		require.NoError(t, gotErr)
		require.Equal(t, cfg.ActionContinue, got)
		require.Equal(t, "Your Pokedex is empty.\n", output.String())
	})

	t.Run("it shows all pokemon caught", func(t *testing.T) {
		t.Parallel()

		var output bytes.Buffer

		config.User.Pokedex.Catch(dto.Pokemon{Name: "zubat"})
		config.User.Pokedex.Catch(dto.Pokemon{Name: "ditto"})

		got, gotErr := commandPokedex(config, &output)

		require.NoError(t, gotErr)
		require.Equal(t, cfg.ActionContinue, got)
		require.Equal(t, "Your Pokedex:\n  - ditto\n  - zubat\n", output.String())
	})
}
