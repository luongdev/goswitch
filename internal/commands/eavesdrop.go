package commands

import (
	"github.com/luongdev/goswitch/types"
	"github.com/luongdev/fsgo/command/call"
)

type EavesdropCommand struct {
	UId
}

func (a *EavesdropCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	return (&call.Execute{UUID: a.uid, AppName: "eavesdrop"}).BuildMessage(), nil
}

func NewEavesdropCommand(uid string) *EavesdropCommand {
	return &EavesdropCommand{UId: UId{uid: uid}}
}

var _ types.Command = (*EavesdropCommand)(nil)
