package valkey

import (
	"auth-proxy/share"
	"context"
	"strconv"

	"github.com/valkey-io/valkey-go"
)

func SetSession(client valkey.Client, sessionId string, id int) {
	ctx := context.Background()
	client.Do(ctx, client.B().Setex().Key(sessionId).Seconds(share.ExpireSessionTime).Value(strconv.Itoa(id)).Build())
}
