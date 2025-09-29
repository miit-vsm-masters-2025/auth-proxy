package main

import (
	"context"
	"fmt"

	"github.com/valkey-io/valkey-go"
)

func test() {
	ValKeyClient, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})

	if err != nil {
		panic(err)
	}
	defer ValKeyClient.Close()

	ctx := context.Background()
	result := ValKeyClient.Do(ctx, ValKeyClient.B().Set().Key("test").Value("val").Nx().Build())
	fmt.Println(result.AsStrMap())
	result = ValKeyClient.Do(ctx, ValKeyClient.B().Get().Key("test").Build())
	str, _ := result.ToString()
	fmt.Println(str)
}
