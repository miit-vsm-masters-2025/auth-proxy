package valkey

import (
	"auth-proxy/share"
	"context"
	"strconv"
	"time"

	"github.com/valkey-io/valkey-go"
)

func SetSession(valkeyClient *valkey.Client, sessionId string, id int) {
	client := *valkeyClient
	ctx := context.Background()
	client.Do(ctx, client.B().Setex().Key(sessionId).Seconds(share.ExpireSessionTime).Value(strconv.Itoa(id)).Build())
}

func CheckSession(valkeyClient *valkey.Client, sessionId string) (userId int, err error) {
	client := *valkeyClient
	ctx := context.Background()
	userIdStr, err := client.Do(ctx, client.B().Getex().Key(sessionId).Ex(time.Duration(share.ExpireSessionTime)).Build()).ToString()
	_, err = strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
