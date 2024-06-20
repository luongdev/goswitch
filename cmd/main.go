package main

import (
	"context"
	"github.com/luongdev/goswitch"
	"github.com/luongdev/goswitch/pkg"
	"time"
)

func main() {
	ic := &goswitch.InboundConfig{
		Host:           "103.141.141.57",
		Port:           65021,
		Password:       "Simplefs!!",
		ConnectTimeout: 10,
	}

	c, err := ic.Build()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := c.Exec(ctx, pkg.ShowCommand("status"))
	if err != nil {
		panic(err)
	}

	if res.IsOk() {
		println(res.GetBody())
	}

	defer c.Disconnect()
}
