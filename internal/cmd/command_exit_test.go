package cmd

import (
	"bytes"
	"testing"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/stretchr/testify/require"
)

func Test_commandExit(t *testing.T) {
	t.Run("it returns ActionExit", func(t *testing.T) {
		t.Parallel()

		config := &cfg.Config{Commands: Commands()}

		var output bytes.Buffer

		got, gotErr := commandExit(config, &output)

		require.NoError(t, gotErr)
		require.Equal(t, cfg.ActionExit, got)
		require.Equal(t, "Closing the Pokedex... Goodbye!\n", output.String())
	})
}
