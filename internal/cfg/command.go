package cfg

import (
	"io"
)

type Command struct {
	Name        string
	Signature   string
	Description string
	Callback    func(config *Config, out io.Writer, args ...string) (Action, error)
}
