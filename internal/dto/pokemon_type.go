package dto

type PokemonType struct {
	// The order the Pokémon's types are listed in.
	Slot int `json:"slot"`

	// The type the referenced Pokémon has.
	Type NamedAPIResource[Type] `json:"type"`
}
