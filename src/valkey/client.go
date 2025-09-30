package valkey

import (
	"github.com/valkey-io/valkey-go"
)

func Client(address string) *valkey.Client {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{address}})
	if err != nil {
		panic(err)
	}

	return &client
}
