package routes

import (
	"auth-proxy/postgre"
	"auth-proxy/share"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func parseRegisterBody(ctx *gin.Context) (*share.Reg, error) {
	var reg share.Reg
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &reg)
	return &reg, err
}

func Register(appContext *share.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		registerForm, err := parseRegisterBody(ctx)
		if err != nil {
			appContext.Logger.Debugf(err.Error())
			ctx.AbortWithStatus(400)
			return
		}
		password, err := bcrypt.GenerateFromPassword([]byte(registerForm.Password), 0)
		login := registerForm.Login
		if err != nil {
			panic(err)
		}
		err = postgre.AddUser(
			appContext.PostgresClient,
			login,
			string(password),
			"",
		)
		if err != nil {
			appContext.Logger.Warnf(err.Error())
			ctx.AbortWithStatus(500)
			return
		}

	}

}
