package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/valkey-io/valkey-go"
)

func test() {
	ValKeyClient, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})

	if err != nil {
		panic(err)
	}
	defer ValKeyClient.Close()

	ctx := context.Background()
	result := ValKeyClient.Do(ctx, ValKeyClient.B().Set().Key("test").Value(strconv.Itoa(99)).Build())
	fmt.Println(result)
	fmt.Println(ValKeyClient.Do(ctx, ValKeyClient.B().Get().Key("test").Build()).ToInt64())
	//var str string
	//str, err = ValKeyClient.Do(ctx, ValKeyClient.B().Get().Key("test").Build()).ToInt64()
	//r, _ := strconv.ParseInt(str, 10, 64)
	//fmt.Println(err)
	//fmt.Println(str)
	//fmt.Println(r)
}
