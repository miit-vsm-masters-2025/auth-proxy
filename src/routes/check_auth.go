package routes

import "github.com/gin-gonic/gin"

func CheckAuth(c *gin.Context) {
	sessionId := c.Request.Header.Get("Cookie")
	if sessionId == "" {
		c.AbortWithStatus(401)
		return
	}

}
