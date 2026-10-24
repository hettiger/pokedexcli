package tests

import (
	"testing"

	"github.com/hettiger/pokedexcli/internal/dto"
	"github.com/hettiger/pokedexcli/internal/svc"
	"github.com/stretchr/testify/require"
)

type PokeAPIFake struct {
	T       *testing.T
	Call    int
	Results []PokeAPIFakeCallResult
}

type PokeAPIFakeCallResult interface {
	isPokeAPIFakeCallResult()
}

func (l LocationAreasFakeResult) isPokeAPIFakeCallResult() {}
func (l LocationAreaFakeResult) isPokeAPIFakeCallResult()  {}

type LocationAreasFakeResult struct {
	Direction svc.Direction
	Areas     dto.APIResourceList[dto.LocationArea]
	Err       error
}

type LocationAreaFakeResult struct {
	Name string
	Area dto.LocationArea
	Err  error
}

func (p *PokeAPIFake) LocationAreas(direction svc.Direction) (dto.APIResourceList[dto.LocationArea], error) {
	p.T.Helper()
	require.Less(p.T, p.Call, len(p.Results), "unexpected LocationAreas call")
	result, ok := p.Results[p.Call].(LocationAreasFakeResult)
	require.True(p.T, ok, "call %d: expected LocationAreasFakeResult, got %T", p.Call, p.Results[p.Call])
	require.Equal(p.T, result.Direction, direction)
	p.Call++
	return result.Areas, result.Err
}

func (p *PokeAPIFake) LocationArea(name string) (dto.LocationArea, error) {
	p.T.Helper()
	require.Less(p.T, p.Call, len(p.Results), "unexpected LocationArea call")
	result, ok := p.Results[p.Call].(LocationAreaFakeResult)
	require.True(p.T, ok, "call %d: expected LocationAreaFakeResult, got %T", p.Call, p.Results[p.Call])
	require.Equal(p.T, result.Name, name)
	p.Call++
	return result.Area, result.Err
}
