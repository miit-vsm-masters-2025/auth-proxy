package main

import (
	"context"
	"fmt"

	"github.com/valkey-io/valkey-go"
)

func test() {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{"localhost:6379"}})
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	msg := client.Do(ctx, client.B().Get().Key("test").Build())
	fmt.Println(msg.ToString())
}
