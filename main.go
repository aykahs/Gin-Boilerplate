package main

import (
	"github.com/aykahs/Gin-Boilerplate/configs"
	"github.com/aykahs/Gin-Boilerplate/pkg/mysql"
	"github.com/aykahs/Gin-Boilerplate/router"
	"github.com/gin-gonic/gin"

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
	configs.Init()
	EnvConfig := configs.EnvConfig
	mysql.Connect(&EnvConfig.Mysql)
	// gin.SetMode(gin.DebugMode)
	router.Init()
	r := router.Router
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Run(EnvConfig.Server.Port)
}
