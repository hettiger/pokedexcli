package user

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pokedex(t *testing.T) {
	t.Run("it stores caught pokemon", func(t *testing.T) {
		t.Parallel()

		p := newPokedex()

		require.Empty(t, p.Caught())

		p.Catch("pikachu")

		require.Equal(t, []string{"pikachu"}, p.Caught())

		p.Catch("mew")

		require.Equal(t, []string{"mew", "pikachu"}, p.Caught())
	})
}
