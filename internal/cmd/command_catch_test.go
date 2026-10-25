package cmd

import (
	"bytes"
	"testing"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/tests"
	"github.com/stretchr/testify/require"
)

func Test_commandCatch(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		results []tests.PokeAPIFakeCallResult
		want    string
		wantErr error
	}{
		{
			name:    "missing arg",
			wantErr: ErrCommandCatchMissingPokemonName,
		},
		{
			name: "missing pokemon",
			args: []string{"fake-pokemon"},
			results: []tests.PokeAPIFakeCallResult{
				tests.PokemonFakeResult{
					Name: "fake-pokemon",
					Err:  tests.ErrFake404,
				},
			},
			wantErr: tests.ErrFake404,
		},
		{
			name: "existing pokemon, caught",
			args: []string{"fake-pokemon"},
			results: []tests.PokeAPIFakeCallResult{
				tests.PokemonFakeResult{
					Name: "fake-pokemon",
					Pokemon: dto.Pokemon{
						Name:           "fake-pokemon",
						BaseExperience: 36,
					},
				},
			},
			want: "Throwing a Pokeball at fake-pokemon...\nChance: 90.00%\nfake-pokemon was caught!\n",
		},
		{
			name: "existing pokemon, not caught",
			args: []string{"fake-pokemon"},
			results: []tests.PokeAPIFakeCallResult{
				tests.PokemonFakeResult{
					Name: "fake-pokemon",
					Pokemon: dto.Pokemon{
						Name:           "fake-pokemon",
						BaseExperience: 608,
					},
				},
			},
			want: "Throwing a Pokeball at fake-pokemon...\nChance: 5.00%\nfake-pokemon escaped!\n",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			config := tests.NewConfigFake(Commands(), tests.NewPokeAPIFake(t, tt.results))

			var output bytes.Buffer

			got, gotErr := commandCatch(config, &output, tt.args...)

			require.ErrorIs(t, tt.wantErr, gotErr)
			require.Equal(t, cfg.ActionContinue, got)
			require.Equal(t, tt.want, output.String())
		})
	}
}
