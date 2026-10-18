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

func Test_commandMapb(t *testing.T) {
	cases := []struct {
		name    string
		results []tests.LocationAreasFakeResult
		want    string
		wantErr error
	}{
		{
			name: "empty",
			results: []tests.LocationAreasFakeResult{
				{Direction: svc.Backward},
			},
			want: "you're on the first page\n",
		},
		{
			name: "two pages",
			results: []tests.LocationAreasFakeResult{
				{
					Direction: svc.Backward,
					Areas: dto.APIResourceList[dto.LocationArea]{
						Results: []dto.LocationArea{
							{Name: "baz"},
						},
					},
				},
				{
					Direction: svc.Backward,
					Areas: dto.APIResourceList[dto.LocationArea]{
						Results: []dto.LocationArea{
							{Name: "foo"},
							{Name: "bar"},
						},
					},
				},
				{Direction: svc.Backward},
			},
			want: "baz\nfoo\nbar\nyou're on the first page\n",
		},
		{
			name: "error",
			results: []tests.LocationAreasFakeResult{
				{
					Direction: svc.Backward,
					Err:       tests.ErrFakeAPI,
				},
			},
			wantErr: tests.ErrFakeAPI,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			config := &cfg.Config{
				Commands: Commands(),
				PokeAPI: &tests.LocationAreasFake{
					T:       t,
					Results: tt.results,
				},
			}

			var output bytes.Buffer

			for i := 0; i < len(tt.results); i++ {
				got, gotErr := commandMapb(config, &output)

				require.ErrorIs(t, gotErr, tt.wantErr)
				require.Equal(t, cfg.ActionContinue, got)
			}

			require.Equal(t, tt.want, output.String())
		})
	}
}
