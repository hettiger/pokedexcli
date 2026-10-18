package cfg

import (
	"io"
)

type Command struct {
	Name        string
	Description string
	Callback    func(config *Config, out io.Writer) (Action, error)
}
