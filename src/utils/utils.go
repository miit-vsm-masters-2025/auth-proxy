package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoadEnv(logger *zap.SugaredLogger) {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
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

func GetPostgresConf() (host, port, user, password, dbname string) {
	host = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("POSTGRES_PORT")
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DB")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}
	return host, port, user, password, dbname
}

func Logger() (*zap.SugaredLogger, func()) {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	return sugar, func() { logger.Sync() }
}
