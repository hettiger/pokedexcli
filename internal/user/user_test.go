package user_test

import (
	"testing"

	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/user"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	t.Run("it provides access to the pokedex", func(t *testing.T) {
		got := user.NewUser()

		got.Pokedex.Catch(dto.Pokemon{Name: "mew"})

		require.Equal(t, []dto.Pokemon{{Name: "mew"}}, got.Pokedex.Caught())
	})
}
