package routes

import (
	"auth-proxy/share"

	"github.com/gin-gonic/gin"
)

func Ping(appContext *share.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	}
}
