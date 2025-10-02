package routes

import (
	"auth-proxy/postgre"
	"auth-proxy/share"
	"auth-proxy/utils"
	localvalkey "auth-proxy/valkey"
	"net/http"
	"strings"

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
	localvalkey.SetSession(client, sessionId, userId)
	return sessionId
}

func parseLoginBody(ctx *gin.Context) *share.Login {
	return &share.Login{
		Login:    strings.TrimSpace(ctx.PostForm("login")),
		Password: ctx.PostForm("password"),
	}
}

func renderLoginFail(ctx *gin.Context, redirectUrl string) {
	if redirectUrl == "" {
		redirectUrl = "/"
	}
	htmlPage := utils.RenderSignInHTML(redirectUrl, "", "")
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Status(http.StatusOK)
	_, _ = ctx.Writer.Write([]byte(htmlPage))
}

func Login(appContext *share.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginForm := parseLoginBody(ctx)
		redirect := ctx.PostForm("return_to")
		id, passwordHash, err := postgre.GetUserId(appContext.PostgresClient, loginForm.Login)
		if err != nil {
			appContext.Logger.Debugf("Login fail: incorrect creditinals %s", loginForm.Login)
			renderLoginFail(ctx, redirect)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginForm.Password))
		if err != nil {
			appContext.Logger.Debugf("Login fail: incorrect password %s", loginForm.Login)
			renderLoginFail(ctx, redirect)
			return
		}
		sessionId := saveSession(appContext.Valkey, id)
		// TODO think about make it secure (only ssl/TLS cookie)
		ctx.Header("X-Session-Id", sessionId)
		ctx.SetCookie("koty42-session-id", sessionId, 3600, "", ".koty42.ru", false, true)
		ctx.Status(200)
		ctx.Redirect(http.StatusSeeOther, redirect)
		return
	}
}
