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

func parseLoginBody(ctx *gin.Context) (*share.Login, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	var loginForm share.Login
	err = json.Unmarshal(body, &loginForm)
	if err != nil {
		return nil, err
	}
	return &loginForm, nil
}

func Login(appContext *share.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginForm, err := parseLoginBody(ctx)
		if err != nil {
			appContext.Logger.Error(err.Error())
			ctx.AbortWithStatus(400)
			return
		}
		id, passwordHash, err := postgre.GetUserId(appContext.PostgresClient, loginForm.Login)
		if err != nil {
			appContext.Logger.Debugf("Login fail: incorrect creditinals %s", loginForm.Login)
			ctx.AbortWithStatus(401)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginForm.Password))
		if err != nil {
			appContext.Logger.Debugf("Login fail: incorrect password %s", loginForm.Login)
			ctx.AbortWithStatus(401)
			return
		}
		sessionId := saveSession(appContext.Valkey, id)
		// TODO think about make it secure (only ssl/TLS cookie)
		ctx.Header("X-Session-Id", sessionId)
		ctx.SetCookie("koty42-session-id", sessionId, 3600, "", ".koty42.ru", false, true)
		ctx.Status(200)
		return
	}
}
