package main

import (
	"github.com/gin-gonic/gin"

	"auth-proxy/routes"
	"auth-proxy/utils"
	"auth-proxy/valkey"
	"go.uber.org/zap"
)

func createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", routes.Ping)
	router.Any("/_auth", routes.CheckAuth)
	return router
}

//type struct AppContext {
//	valkey string
//	logger *zap.Logger
//}

func main() {
	loggerHandler, sync := utils.Logger()
	defer sync()

	utils.LoadEnv()
	router := createRouter()
	// that is bad practise https://stackoverflow.com/questions/35672842/go-and-gin-passing-around-struct-for-database-context
	router.Use(valkey.ValkeyClient(utils.GetValKeyAddress())) // create middleware that provide valkey client
	router.Use(loggerHandler)                                 // create middleware that provide logger
	router.Run(":8080")
}
