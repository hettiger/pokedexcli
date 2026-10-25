package dto_test

import (
	"fmt"
	"testing"

	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/stretchr/testify/require"
)

func TestPokemon_Chance(t *testing.T) {
	tests := []struct {
		xp   int
		want float64
	}{
		{xp: -10, want: 90},
		{xp: 0, want: 90},
		{xp: 1, want: 90},
		{xp: 32, want: 90},
		{xp: 100, want: 65.13},
		{xp: 400, want: 14.3},
		{xp: 608, want: 5},
		{xp: 700, want: 5},
		{xp: 9999, want: 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.xp), func(t *testing.T) {
			p := dto.Pokemon{BaseExperience: tt.xp}
			got := p.Chance()
			require.InDelta(t, tt.want, got, 0.1)
		})
	}
}
