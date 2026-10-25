package tests

import (
	"errors"
	"fmt"

	"github.com/hettiger/pokedexcli/internal/svc"
)

var (
	ErrFakeAPI = errors.New("fake api error")
	ErrFake404 = fmt.Errorf("%w: status %d", svc.ErrNonSuccessStatus, 404)
)
