package valkey

import (
	"github.com/gin-gonic/gin"
	"github.com/valkey-io/valkey-go"
)

func ValkeyClient(address string) gin.HandlerFunc {

	ValKeyClient, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{address}})
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set("valkey", ValKeyClient)
		c.Next()
	}
}
