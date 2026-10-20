package svc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/hettiger/pokedexcli/internal/dto"
)

type PokeAPI struct {
	baseURL               string
	cache                 cacher
	nextLocationAreas     *string
	previousLocationAreas *string
}

type Direction int

const (
	Forward Direction = iota
	Backward
)

type cacher interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
}

func NewPokeAPI(baseURL string, cache cacher) *PokeAPI {
	return &PokeAPI{
		baseURL: baseURL,
		cache:   cache,
	}
}

var ErrNonSuccessStatus = errors.New("non-successful PokéAPI response")

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

	data, ok := p.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return dto.APIResourceList[dto.LocationArea]{}, err
		}
		defer res.Body.Close()

		if res.StatusCode >= 300 {
			err := fmt.Errorf("%w: status %d", ErrNonSuccessStatus, res.StatusCode)
			return dto.APIResourceList[dto.LocationArea]{}, err
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return dto.APIResourceList[dto.LocationArea]{}, err
		}

		p.cache.Add(url, data)
	}

	areas := dto.APIResourceList[dto.LocationArea]{}
	if err := json.Unmarshal(data, &areas); err != nil {
		return dto.APIResourceList[dto.LocationArea]{}, err
	}

	p.nextLocationAreas = areas.Next
	p.previousLocationAreas = areas.Previous

	return areas, nil
}
