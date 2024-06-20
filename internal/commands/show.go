package commands

import (
	"fmt"
	"github.com/luongdev/goswitch/types"
	"github.com/percipia/eslgo/command"
)

type ShowCommand struct {
	Obj string
}

func (a *ShowCommand) Validate() error {
	if a.Obj == "" {
		return fmt.Errorf("object must be provided")
	}

	return nil
}

func (a *ShowCommand) Raw() (string, error) {

	return (&command.API{Command: "show", Arguments: a.Obj}).BuildMessage(), nil
}

func NewShowCommand(obj string) *ShowCommand {
	return &ShowCommand{Obj: obj}
}

var _ types.Command = (*ShowCommand)(nil)
