package main

import (
	"auth-proxy/utils"
	valkeypackage "auth-proxy/valkey"

	"github.com/gin-gonic/gin"

	"auth-proxy/postgre"
	"auth-proxy/routes"
	"auth-proxy/share"
)

func createRouter(appCtx share.AppContext) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", routes.Ping(&appCtx))
	router.Any("/_auth", routes.CheckAuth(&appCtx))
	auth := router.Group("/user")
	{
		auth.GET("/me", routes.CheckAuth(&appCtx))
		auth.GET("/login", routes.CheckAuth(&appCtx))
		auth.GET("/reg", routes.CheckAuth(&appCtx))
	}
	router.POST("/user/login", routes.CheckAuth(&appCtx))
	return router
}

func main() {
	logger, sync := utils.Logger()
	defer sync()

	utils.LoadEnv(logger)

	valkeyClient := valkeypackage.ValkeyClient(utils.GetValKeyAddress())
	defer (*valkeyClient).Close()

	postgresClient := postgre.PostgresClient(utils.GetPostgresConf())
	defer (*postgresClient).Close()
	postgre.CreateTable(postgresClient)

	appCtx := share.AppContext{Valkey: valkeyClient, Logger: logger, PostgresClient: postgresClient}
	router := createRouter(appCtx)
	router.Run(":8080")
}
