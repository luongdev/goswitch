package internal

import (
	"fmt"
	"github.com/luongdev/fsgo"
	"github.com/luongdev/goswitch/types"
)

type eventImpl struct {
	raw *eslgo.Event
}

func (e *eventImpl) HasHeader(name string) bool {
	return e.raw.HasHeader(name)
}

func (e *eventImpl) GetHeader(name string) string {
	return e.raw.GetHeader(name)
}

func (e *eventImpl) GetVariable(name string) string {
	return e.raw.GetHeader(fmt.Sprintf("variable_%v", name))
}

func NewEvent(raw *eslgo.Event) types.Event {
	return &eventImpl{raw: raw}
}

var _ types.Event = (*eventImpl)(nil)
