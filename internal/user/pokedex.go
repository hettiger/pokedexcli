package user

import (
	"cmp"
	"slices"

	"github.com/hettiger/pokedexcli/internal/dto"
)

type pokedex struct {
	caught map[string]dto.Pokemon
}

func newPokedex() pokedex {
	return pokedex{
		caught: make(map[string]dto.Pokemon),
	}
}

// Catch adds the given Pokémon to the caught set
func (p *pokedex) Catch(pokemon dto.Pokemon) {
	p.caught[pokemon.Name] = pokemon
}

func (p *pokedex) Inspect(name string) (dto.Pokemon, bool) {
	pokemon, ok := p.caught[name]
	if !ok {
		return dto.Pokemon{}, false
	}

	return pokemon, true
}

// Caught returns all Pokémon caught in sorted order.
func (p *pokedex) Caught() []dto.Pokemon {
	pokemon := make([]dto.Pokemon, 0, len(p.caught))

	for _, p := range p.caught {
		pokemon = append(pokemon, p)
	}

	slices.SortFunc(pokemon, func(a, b dto.Pokemon) int {
		return cmp.Compare(a.Name, b.Name)
	})

	return pokemon
}
