package routes

import (
	"auth-proxy/share"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

package routes

import (
"auth-proxy/share"
"strconv"

"github.com/gin-gonic/gin"
)

func Login(appContext *share.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, _ :=  io.ReadAll(ctx.Request.Body)
		var loginForm share.Login
		if err := json.Unmarshal(body, &loginForm); err != nil {
			appContext.Logger.Error(err.Error())
			ctx.AbortWithStatus(400)
		}
		
	}
}
