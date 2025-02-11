package commands

import (
	"github.com/luongdev/fsgo/command/call"
	"github.com/luongdev/goswitch/types"
)

type AnswerCommand struct {
	UId
}

func (a *AnswerCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	return (&call.Execute{UUID: a.uid, AppName: "answer"}).BuildMessage(), nil
}

func NewAnswerCommand(uid string) *AnswerCommand {
	return &AnswerCommand{UId: UId{uid: uid}}
}

var _ types.Command = (*AnswerCommand)(nil)
