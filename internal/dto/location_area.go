package dto

// LocationArea is a partial representation of the original PokéAPI type.
type LocationArea struct {
	// The identifier for this resource.
	Id int `json:"id"`

	// The name for this resource.
	Name string `json:"name"`

	// A list of Pokémon that can be encountered in this area.
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
