package dto

import "math"

const (
	minBaseXP float64 = 36
	maxBaseXP float64 = 608
)

// Pokemon is a partial representation of the original PokéAPI type.
type Pokemon struct {
	// The identifier for this resource.
	Id int `json:"id"`

	// The name for this resource.
	Name string `json:"name"`

	// The base experience gained for defeating this Pokémon.
	BaseExperience int `json:"base_experience"`
}

// Chance returns a catch chance from 5% to 90% using exponential decay.
func (p *Pokemon) Chance() float64 {
	xp := float64(p.BaseExperience)
	chance := 90 * math.Pow(90.0/5.0, -(xp-minBaseXP)/(maxBaseXP-minBaseXP))
	return math.Max(5, math.Min(90, chance))
}
