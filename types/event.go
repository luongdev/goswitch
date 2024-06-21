package types

import "context"

type Event interface {
	HasHeader(name string) bool
	GetHeader(name string) string
	GetVariable(name string) string
}

type EventHandler interface {
	Handle(ctx context.Context, e Event)
}
