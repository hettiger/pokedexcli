package cmd

import (
	"bytes"
	"testing"

	"github.com/hettiger/pokedexcli/internal/cfg"
	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/svc"
	"github.com/hettiger/pokedexcli/internal/tests"
	"github.com/stretchr/testify/require"
)

func Test_commandMap(t *testing.T) {
	cases := []struct {
		name    string
		results []tests.PokeAPIFakeCallResult
		want    string
		wantErr error
	}{
		{
			name: "empty",
			results: []tests.PokeAPIFakeCallResult{
				tests.LocationAreasFakeResult{
					Direction: svc.Forward,
				},
			},
			want: "you're on the last page\n",
		},
		{
			name: "two pages",
			results: []tests.PokeAPIFakeCallResult{
				tests.LocationAreasFakeResult{
					Direction: svc.Forward,
					Areas: dto.APIResourceList[dto.LocationArea]{
						Results: []dto.LocationArea{
							{Name: "foo"},
							{Name: "bar"},
						},
					},
				},
				tests.LocationAreasFakeResult{
					Direction: svc.Forward,
					Areas: dto.APIResourceList[dto.LocationArea]{
						Results: []dto.LocationArea{
							{Name: "baz"},
						},
					},
				},
				tests.LocationAreasFakeResult{
					Direction: svc.Forward,
				},
			},
			want: "foo\nbar\nbaz\nyou're on the last page\n",
		},
		{
			name: "error",
			results: []tests.PokeAPIFakeCallResult{
				tests.LocationAreasFakeResult{
					Direction: svc.Forward,
					Err:       tests.ErrFakeAPI,
				},
			},
			wantErr: tests.ErrFakeAPI,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			config := tests.NewConfigFake(Commands(), tests.NewPokeAPIFake(t, tt.results))

			var output bytes.Buffer

			for i := 0; i < len(tt.results); i++ {
				got, gotErr := commandMap(config, &output)

				require.ErrorIs(t, gotErr, tt.wantErr)
				require.Equal(t, cfg.ActionContinue, got)
			}

			require.Equal(t, tt.want, output.String())
		})
	}
}
