package main

import (
	"fmt"
	"os"

	"github.com/aykahs/Gin-Boilerplate/configs"
	"github.com/aykahs/Gin-Boilerplate/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	uuid "github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	env := os.Getenv("APP_ENV")
	ginmode := os.Getenv("GIN_MODE")

	configs.Init(env)
	EnvConfig := configs.EnvConfig
	// mysql.Connect(&EnvConfig.Mysql)
	gin.SetMode(ginmode)
	router.Init()
	r := router.Router
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Run(EnvConfig.Server.Port)
}
