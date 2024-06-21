package types

import (
	"context"
)

type Client interface {
	Disconnect()
	Exec(ctx context.Context, cmd Command) (CommandOutput, error)
	Events(ctx context.Context, events ...string) error
	AddEventHandler(key string, handler EventHandler)
	GetSessionId() string
}

type ClientDisconnectFunc func(ctx context.Context)
