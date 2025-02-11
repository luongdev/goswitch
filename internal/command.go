package internal

import (
	"github.com/luongdev/fsgo"
	"github.com/luongdev/goswitch/types"
	"strings"
)

type CommandImpl struct {
	cmd string
}

func (c *CommandImpl) BuildMessage() string {
	return c.cmd
}

func NewCommand(cmd string) *CommandImpl {
	return &CommandImpl{cmd: cmd}
}

type CommandOutputImpl struct {
	*eslgo.RawResponse
}

func (c *CommandOutputImpl) GetBody() string {
	if c.RawResponse == nil || c.RawResponse.Body == nil {
		return ""
	}

	return string(c.RawResponse.Body)
}

func (c *CommandOutputImpl) GetReply() string {
	return c.RawResponse.GetReply()
}

func (c *CommandOutputImpl) IsOk() bool {
	reply := c.GetReply()
	if reply == "" {
		return false
	}

	if strings.HasPrefix(reply, "-ERR") {
		return false
	}

	return true
}

func NewCommandOutput(raw *eslgo.RawResponse) *CommandOutputImpl {
	return &CommandOutputImpl{RawResponse: raw}
}

var _ types.CommandOutput = (*CommandOutputImpl)(nil)
