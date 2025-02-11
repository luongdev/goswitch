package commands

import (
	"github.com/luongdev/fsgo/command/call"
	"github.com/luongdev/goswitch/types"
)

type EchoCommand struct {
	UId
}

func (a *EchoCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	return (&call.Execute{UUID: a.uid, AppName: "echo"}).BuildMessage(), nil
}

func NewEchoCommand(uid string) *EchoCommand {
	return &EchoCommand{UId: UId{uid: uid}}
}

var _ types.Command = (*EchoCommand)(nil)
