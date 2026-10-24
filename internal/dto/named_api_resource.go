package dto

type NamedAPIResource[TType any] struct {
	// The name of the referenced resource.
	Name string `json:"name"`

	// The URL of the referenced resource.
	URL string `json:"url"`
}
