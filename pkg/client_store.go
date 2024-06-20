package pkg

import (
	"github.com/luongdev/goswitch/internal"
	"github.com/luongdev/goswitch/types"
)

func NewClientStore(m map[string]types.Client) types.ClientStore {
	return internal.NewClientStore(m)
}
