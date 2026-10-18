package dto

// APIResourceList provides pagination metadata for a PokéAPI resources list.
type APIResourceList[TResource any] struct {
	// The total number of resources available from this API.
	Count int `json:"count"`

	// The URL for the next page in the list.
	Next *string `json:"next"`

	// The URL for the previous page in the list.
	Previous *string `json:"previous"`

	// A list of unnamed API resources.
	Results []TResource `json:"results"`
}
