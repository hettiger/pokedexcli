package dto

// PokemonEncounter is a partial representation of the original PokéAPI type.
type PokemonEncounter struct {
	// The Pokémon being encountered.
	Pokemon NamedAPIResource[Pokemon] `json:"pokemon"`
}
