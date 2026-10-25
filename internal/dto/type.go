package dto

// Type is a partial representation of the original PokéAPI type.
type Type struct {
	// The identifier for this resource.
	Id int `json:"id"`

	// The name for this resource.
	Name string `json:"name"`
}
