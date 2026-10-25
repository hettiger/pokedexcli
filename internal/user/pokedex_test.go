package user

import (
	"testing"

	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/stretchr/testify/require"
)

func Test_pokedex(t *testing.T) {
	t.Run("it stores caught pokemon", func(t *testing.T) {
		t.Parallel()

		p := newPokedex()

		require.Empty(t, p.Caught())

		p.Catch(dto.Pokemon{Name: "pikachu"})

		require.Equal(t, []dto.Pokemon{{Name: "pikachu"}}, p.Caught())

		p.Catch(dto.Pokemon{Name: "mew"})

		require.Equal(t, []dto.Pokemon{{Name: "mew"}, {Name: "pikachu"}}, p.Caught())
	})

	t.Run("it allows inspecting a pokemon by name", func(t *testing.T) {
		t.Parallel()

		p := newPokedex()

		pokemon := dto.Pokemon{Name: "ditto"}

		p.Catch(pokemon)

		got, ok := p.Inspect("ditto")

		require.True(t, ok)
		require.Equal(t, pokemon, got)
	})

	t.Run("Inspect returns false if the given pokemon has not been caught yet", func(t *testing.T) {
		t.Parallel()

		p := newPokedex()

		got, ok := p.Inspect("ditto")

		require.False(t, ok)
		require.Equal(t, dto.Pokemon{}, got)
	})
}
