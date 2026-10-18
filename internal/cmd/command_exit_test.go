package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_commandExit(t *testing.T) {
	t.Run("it returns ActionExit", func(t *testing.T) {
		t.Parallel()

		var output bytes.Buffer

		got, gotErr := commandExit(&output)

		require.NoError(t, gotErr)
		require.Equal(t, ActionExit, got)
		require.Equal(t, "Closing the Pokedex... Goodbye!\n", output.String())
	})
}
