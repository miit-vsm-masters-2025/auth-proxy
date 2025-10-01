package routes

import (
	"auth-proxy/share"
	valkeypackage "auth-proxy/valkey"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckAuth(appContext *share.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId := ctx.Request.Header.Get("Cookie")
		if sessionId == "" {
			appContext.Logger.Warnf(
				"CheckAuth fail: No session id in Cookie\n User Headers:%v",
				ctx.Request.Header,
			)
			ctx.AbortWithStatus(401)
			return
		}
		userId, err := valkeypackage.CheckSession(appContext.Valkey, sessionId)
		if err != nil {
			appContext.Logger.Errorf(
				"CheckAuth fail: Bad user id '%s' for session %s\n User Headers:%v",
				userId,
				sessionId,
			)
			//TODO remove invalid session if exists
			ctx.Header("Set-Cookie", "")
			ctx.AbortWithStatus(401)
			return
		}
		appContext.Logger.Debugf("CheckAuth sucess userId:%d", userId)
		ctx.Header("X-User-Id", userId)
		ctx.Status(200)
	}
}
