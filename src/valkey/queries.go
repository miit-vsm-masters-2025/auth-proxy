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
	result := client.Do(ctx, client.B().Getex().Key(sessionId).Ex(time.Duration(share.ExpireSessionTime)).Build())
	if result.Error() != nil {
		return 0, err
	}
	userIdStr, err := result.ToString()
	_, err = strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		//logger.Errorf(
		//	"CheckAuth fail: Bad user id '%s' for session %s\n User Headers:%v",
		//	userId,
		//	sessionId,
		//)
		return 0, err
	}
	return userId, nil
}
