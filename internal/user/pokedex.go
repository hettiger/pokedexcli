package user

import "slices"

type pokedex struct {
	caught map[string]struct{}
}

func newPokedex() pokedex {
	return pokedex{
		caught: make(map[string]struct{}),
	}
}

// Catch adds the given Pokémon to the caught set
func (p *pokedex) Catch(pokemon string) {
	p.caught[pokemon] = struct{}{}
}

// Caught returns all Pokémon caught in sorted order.
func (p *pokedex) Caught() []string {
	names := make([]string, 0, len(p.caught))

	for name := range p.caught {
		names = append(names, name)
	}

	slices.Sort(names)

	return names
}
