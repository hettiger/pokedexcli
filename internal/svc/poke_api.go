package svc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hettiger/pokedexcli/internal/dto"
)

type PokeAPI struct {
	baseURL               string
	nextLocationAreas     *string
	previousLocationAreas *string
}

type Direction int

const (
	Forward Direction = iota
	Backward
)

func NewPokeAPI(baseURL string) *PokeAPI {
	return &PokeAPI{
		baseURL: baseURL,
	}
}

func (p *PokeAPI) LocationAreas(direction Direction) (dto.APIResourceList[dto.LocationArea], error) {
	url := p.baseURL + "/location-area"

	switch {
	case direction == Forward && p.nextLocationAreas == nil && p.previousLocationAreas != nil:
		return dto.APIResourceList[dto.LocationArea]{}, nil
	case direction == Forward && p.nextLocationAreas != nil:
		url = *p.nextLocationAreas
	case direction == Backward && p.previousLocationAreas == nil:
		return dto.APIResourceList[dto.LocationArea]{}, nil
	case direction == Backward && p.previousLocationAreas != nil:
		url = *p.previousLocationAreas
	}

	res, err := http.Get(url)
	if err != nil {
		return dto.APIResourceList[dto.LocationArea]{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		err := fmt.Errorf("non-successful PokéAPI response: %s", res.Status)
		return dto.APIResourceList[dto.LocationArea]{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return dto.APIResourceList[dto.LocationArea]{}, err
	}

	areas := dto.APIResourceList[dto.LocationArea]{}
	if err := json.Unmarshal(data, &areas); err != nil {
		return dto.APIResourceList[dto.LocationArea]{}, err
	}

	p.nextLocationAreas = areas.Next
	p.previousLocationAreas = areas.Previous

	return areas, nil
}
