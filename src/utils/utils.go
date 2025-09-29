package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetValKeyAddress runs after LoadEnv call
func GetValKeyAddress() string {
	port, err := strconv.ParseInt(os.Getenv("VALKEY_PORT"), 10, 0)
	host := os.Getenv("VALKEY_HOST")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func Logger() (gin.HandlerFunc, func()) {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	return func(c *gin.Context) {
		c.Set("logger", sugar)
		c.Next()
	}, func() { logger.Sync() }
}
