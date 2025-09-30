package routes

import (
	"auth-proxy/postgre"
	"auth-proxy/share"
	localvalkey "auth-proxy/valkey"
	"encoding/json"
	"io"

	"crypto/rand"

	"github.com/gin-gonic/gin"
	"github.com/valkey-io/valkey-go"
	"golang.org/x/crypto/bcrypt"
)

func saveSession(client *valkey.Client, userId int) (sessionId string) {
	sessionId = rand.Text()
	if len(sessionId) > 255 {
		sessionId = sessionId[:255]
	}
	localvalkey.SetSession(*client, sessionId, userId)
	return sessionId
}

func Login(appContext *share.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, _ := io.ReadAll(ctx.Request.Body)
		var loginForm share.Login
		if err := json.Unmarshal(body, &loginForm); err != nil {
			appContext.Logger.Error(err.Error())
			ctx.AbortWithStatus(400)
			return
		}
		pass_hash, err := bcrypt.GenerateFromPassword([]byte(loginForm.Password), 0)
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		id, err := postgre.GetUserId(appContext.PostgresClient, loginForm.Login, string(pass_hash))
		if err != nil {
			appContext.Logger.Debugf("Login fail: incorrect creditinals %s", loginForm.Login)
			ctx.AbortWithStatus(401)
			return
		}
		sessionId := saveSession(appContext.Valkey, id)
		// TODO think about make it secure (only ssl/TLS cookie)
		ctx.SetCookie("koty42-session-id", sessionId, 3600, "", ".koty42.ru", false, true)
		ctx.Status(200)
		return
	}
}
