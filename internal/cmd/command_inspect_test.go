package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/tests"
	"github.com/stretchr/testify/require"
)

func Test_commandInspect(t *testing.T) {
	t.Run("it fails on missing pokemon name arg", func(t *testing.T) {
		t.Parallel()

		config := tests.NewConfigFake(Commands(), tests.NewPokeAPIFake(t, nil))

		var output bytes.Buffer

		got, gotErr := commandInspect(config, &output)

		require.ErrorIs(t, ErrMissingPokemonName, gotErr)
		require.Equal(t, cfg.ActionContinue, got)
		require.Equal(t, "", output.String())
	})

	t.Run("it warns on uncaught pokemon", func(t *testing.T) {
		t.Parallel()

		config := tests.NewConfigFake(Commands(), tests.NewPokeAPIFake(t, nil))

		var output bytes.Buffer

		got, gotErr := commandInspect(config, &output, "pidgey")

		require.NoError(t, gotErr)
		require.Equal(t, cfg.ActionContinue, got)
		require.Equal(t, "you have not caught that pokemon\n", output.String())
	})

	t.Run("it shows details about a caught pokemon", func(t *testing.T) {
		t.Parallel()

		config := tests.NewConfigFake(Commands(), tests.NewPokeAPIFake(t, nil))

		config.User.Pokedex.Catch(dto.Pokemon{
			Name:   "pidgey",
			Height: 3,
			Weight: 18,
			Stats: []dto.PokemonStat{
				dto.PokemonStat{
					Stat:     dto.NamedAPIResource[dto.Stat]{Name: "hp"},
					BaseStat: 40,
				},
				dto.PokemonStat{
					Stat:     dto.NamedAPIResource[dto.Stat]{Name: "attack"},
					BaseStat: 45,
				},
				dto.PokemonStat{
					Stat:     dto.NamedAPIResource[dto.Stat]{Name: "defense"},
					BaseStat: 40,
				},
				dto.PokemonStat{
					Stat:     dto.NamedAPIResource[dto.Stat]{Name: "special-attack"},
					BaseStat: 35,
				},
				dto.PokemonStat{
					Stat:     dto.NamedAPIResource[dto.Stat]{Name: "special-defense"},
					BaseStat: 35,
				},
				dto.PokemonStat{
					Stat:     dto.NamedAPIResource[dto.Stat]{Name: "speed"},
					BaseStat: 56,
				},
			},
			Types: []dto.PokemonType{
				dto.PokemonType{
					Type: dto.NamedAPIResource[dto.Type]{Name: "normal"},
				},
				dto.PokemonType{
					Type: dto.NamedAPIResource[dto.Type]{Name: "flying"},
				},
			},
		})

		var output bytes.Buffer

		got, gotErr := commandInspect(config, &output, "pidgey")

		pidgey, err := os.ReadFile("../testdata/pidgey.txt")
		require.NoError(t, err)

		require.NoError(t, gotErr)
		require.Equal(t, cfg.ActionContinue, got)
		require.Equal(t, string(pidgey), output.String())
	})
}
