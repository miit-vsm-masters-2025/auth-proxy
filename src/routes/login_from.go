package routes

import (
	"auth-proxy/share"
	"auth-proxy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(appContext *share.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		redirect := c.Query("return_to")
		htmlPage := utils.RenderSignInHTML(redirect, "", "")
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Status(http.StatusOK)
		_, _ = c.Writer.Write([]byte(htmlPage))
	}
}
