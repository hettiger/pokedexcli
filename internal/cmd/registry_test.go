package cmd_test

import (
	"testing"

	cmd "github.com/hettiger/pokedexcli/internal/cmd"
	"github.com/stretchr/testify/require"
)

func TestCommands(t *testing.T) {
	t.Parallel()

	want := []string{"exit", "help"}

	got := cmd.Commands()

	for _, name := range want {
		require.Contains(t, got, name)
	}

	for name, command := range got {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, name, command.Name)
			require.NotEmpty(t, command.Description)
			require.NotNil(t, command.Callback)
		})
	}
}
