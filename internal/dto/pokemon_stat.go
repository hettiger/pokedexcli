package dto

type PokemonStat struct {
	// The stat the Pokémon has.
	Stat NamedAPIResource[Stat] `json:"stat"`

	// The effort points (EV) the Pokémon has in the stat.
	Effort int `json:"effort"`

	// The base value of the stat.
	BaseStat int `json:"base_stat"`
}
