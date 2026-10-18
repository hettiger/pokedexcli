package repl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_cleanInput(t *testing.T) {
	tests := []struct {
		text string
		want []string
	}{
		{text: "hello world", want: []string{"hello", "world"}},
		{text: "  trimmed whitespace ", want: []string{"trimmed", "whitespace"}},
		{text: "Charmander Bulbasaur PIKACHU", want: []string{"charmander", "bulbasaur", "pikachu"}},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			t.Parallel()

			got := cleanInput(tt.text)

			require.Equal(t, tt.want, got)
		})
	}
}
