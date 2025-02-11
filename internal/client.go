package internal

import (
	"context"
	"fmt"
	"github.com/luongdev/fsgo"
	"github.com/luongdev/fsgo/command"
	"github.com/luongdev/goswitch/types"
)

type ClientImpl struct {
	conn          *eslgo.Conn
	sessionId     string
	ctx           context.Context
	eventHandlers map[string]types.EventHandler
}

func (c *ClientImpl) RemoveEventHandler(key, handlerId string) {
	delete(c.eventHandlers, key)
	c.conn.RemoveEventListener(key, handlerId)
}

func (c *ClientImpl) AddEventHandler(key string, handler types.EventHandler) string {
	c.eventHandlers[key] = handler

	return c.conn.RegisterEventListener(key, func(e *eslgo.Event) {
		event := NewEvent(e)
		h, ok := c.eventHandlers["ALL"]
		if ok && h != nil {
			h.Handle(c.ctx, event)
		}

		h, ok = c.eventHandlers[key]
		if ok && h != nil {
			h.Handle(c.ctx, event)
		}
	})
}

func (c *ClientImpl) GetSessionId() string {
	return c.sessionId
}

func (c *ClientImpl) Disconnect() {
	if c.conn != nil {
		c.conn.ExitAndClose()
	}
}

func (c *ClientImpl) Exec(ctx context.Context, cmd types.Command) (types.CommandOutput, error) {
	in, err := cmd.Raw()
	if err != nil {
		return nil, err
	}

	out, err := c.conn.SendCommand(ctx, NewCommand(in))
	if err != nil {
		return nil, err
	}

	return NewCommandOutput(out), nil
}

func (c *ClientImpl) Events(ctx context.Context, events ...string) error {
	var cmd command.Command
	if len(events) == 0 {
		cmd = &command.DisableEvents{}
		c.eventHandlers = make(map[string]types.EventHandler)
	} else {
		cmd = &command.Event{Format: "plain", Listen: events}
	}

	out, err := c.conn.SendCommand(ctx, cmd)
	if err != nil {
		return err
	}

	res := NewCommandOutput(out)
	if !res.IsOk() {
		return fmt.Errorf("failed to listen to events: %s", res.GetBody())
	}

	return nil
}

func NewClient(c *eslgo.Conn, ctx context.Context) *ClientImpl {
	return &ClientImpl{conn: c, ctx: ctx, eventHandlers: make(map[string]types.EventHandler)}
}
func NewSessionClient(c *eslgo.Conn, sessionId string, ctx context.Context) *ClientImpl {
	return &ClientImpl{conn: c, ctx: ctx, sessionId: sessionId}
}

var _ types.Client = (*ClientImpl)(nil)
