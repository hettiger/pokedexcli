package cmd

import (
	"bytes"
	"testing"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/tests"
	"github.com/stretchr/testify/require"
)

func Test_commandExplore(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		results []tests.PokeAPIFakeCallResult
		want    string
		wantErr error
	}{
		{
			name:    "missing arg",
			wantErr: ErrCommandExploreMissingLocation,
		},
		{
			name: "missing area",
			args: []string{"fake-area"},
			results: []tests.PokeAPIFakeCallResult{
				tests.LocationAreaFakeResult{
					Name: "fake-area",
					Err:  tests.ErrFake404,
				},
			},
			wantErr: tests.ErrFake404,
		},
		{
			name: "existing area, no pokemon encounter",
			args: []string{"fake-area"},
			results: []tests.PokeAPIFakeCallResult{
				tests.LocationAreaFakeResult{
					Name: "fake-area",
					Area: dto.LocationArea{
						Id:   1,
						Name: "fake-area",
					},
				},
			},
			want: "Exploring fake-area...\nFound Pokémon:\nNo Pokémon encounters in this area.\n",
		},
		{
			name: "existing area, one pokemon encounter",
			args: []string{"fake-area"},
			results: []tests.PokeAPIFakeCallResult{
				tests.LocationAreaFakeResult{
					Name: "fake-area",
					Area: dto.LocationArea{
						Id:   1,
						Name: "fake-area",
						PokemonEncounters: []dto.PokemonEncounter{
							{
								Pokemon: dto.NamedAPIResource[dto.Pokemon]{
									Name: "ditto",
								},
							},
						},
					},
				},
			},
			want: "Exploring fake-area...\nFound Pokémon:\n- ditto\n",
		},
		{
			name: "existing area, multiple pokemon encounters",
			args: []string{"fake-area"},
			results: []tests.PokeAPIFakeCallResult{
				tests.LocationAreaFakeResult{
					Name: "fake-area",
					Area: dto.LocationArea{
						Id:   1,
						Name: "fake-area",
						PokemonEncounters: []dto.PokemonEncounter{
							{
								Pokemon: dto.NamedAPIResource[dto.Pokemon]{
									Name: "ditto",
								},
							},
							{
								Pokemon: dto.NamedAPIResource[dto.Pokemon]{
									Name: "glumanda",
								},
							},
						},
					},
				},
			},
			want: "Exploring fake-area...\nFound Pokémon:\n- ditto\n- glumanda\n",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			config := tests.NewConfigFake(Commands(), tests.NewPokeAPIFake(t, tt.results))

			var output bytes.Buffer

			got, gotErr := commandExplore(config, &output, tt.args...)

			require.ErrorIs(t, tt.wantErr, gotErr)
			require.Equal(t, cfg.ActionContinue, got)
			require.Equal(t, tt.want, output.String())
		})
	}
}
