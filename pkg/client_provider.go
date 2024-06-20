package pkg

import (
	"github.com/luongdev/goswitch/internal"
	"github.com/luongdev/goswitch/types"
)

func NewClientProvider(store types.ClientStore) types.ClientProvider {
	return internal.NewClientProvider(store)
}
