package cmd

import "io"

type Command struct {
	Name        string
	Description string
	Callback    func(out io.Writer) (Action, error)
}
