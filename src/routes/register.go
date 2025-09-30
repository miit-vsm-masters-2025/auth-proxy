package routes

import (
	"auth-proxy/share"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(appContext *share.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reg share.Reg
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &reg); err != nil {
			appContext.Logger.Debugf(err.Error())
			ctx.AbortWithStatus(400)
			return
		}
		pass, err := bcrypt.GenerateFromPassword([]byte(reg.Password), 0)
		if err != nil {
			panic(err)
		}
		appContext.PostgresClient.Exec(
			"INSERT INTO user (login, password_hash, totp_secret) VALUES (?, ?, ?)",
			reg.Login,
			pass,
			"",
		)

	}

}
