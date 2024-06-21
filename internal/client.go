package internal

import (
	"context"
	"fmt"
	"github.com/luongdev/goswitch/types"
	"github.com/percipia/eslgo"
	"github.com/percipia/eslgo/command"
)

type ClientImpl struct {
	conn          *eslgo.Conn
	sessionId     string
	ctx           context.Context
	eventHandlers map[string]types.EventHandler
}

func (c *ClientImpl) AddEventHandler(key string, handler types.EventHandler) {
	c.eventHandlers[key] = handler
}

func (c *ClientImpl) RegisterEvents(ctx context.Context, event string) string {
	return c.conn.RegisterEventListener(event, func(raw *eslgo.Event) {
		h, ok := c.eventHandlers["ALL"]
		if ok && h != nil {
			h.Handle(ctx, NewEvent(raw))
		}

		h, ok = c.eventHandlers[event]
		if ok && h != nil {
			h.Handle(ctx, NewEvent(raw))
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

	c.RegisterEvents(ctx, "ALL")
	for _, event := range events {
		c.RegisterEvents(ctx, event)
		if "ALL" == event {
			continue
		}
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
