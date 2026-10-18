package tests

import (
	"testing"

	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/svc"
	"github.com/stretchr/testify/require"
)

type LocationAreasFake struct {
	T       *testing.T
	Call    int
	Results []LocationAreasFakeResult
}

type LocationAreasFakeResult struct {
	Direction svc.Direction
	Areas     dto.APIResourceList[dto.LocationArea]
	Err       error
}

func (p *LocationAreasFake) LocationAreas(direction svc.Direction) (dto.APIResourceList[dto.LocationArea], error) {
	p.T.Helper()
	require.Less(p.T, p.Call, len(p.Results), "unexpected LocationAreas call")
	result := p.Results[p.Call]
	require.Equal(p.T, result.Direction, direction)
	p.Call++
	return result.Areas, result.Err
}
