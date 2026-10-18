package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/stretchr/testify/require"
)

func Test_commandHelp(t *testing.T) {
	t.Run("it returns expected output", func(t *testing.T) {
		t.Parallel()

		config := &cfg.Config{Commands: Commands()}

		var output bytes.Buffer

		want, err := os.ReadFile("../testdata/help.txt")
		require.NoError(t, err)

		got, gotErr := commandHelp(config, &output)

		require.NoError(t, gotErr)
		require.Equal(t, cfg.ActionContinue, got)
		require.Equal(t, string(want), output.String())
	})
}
