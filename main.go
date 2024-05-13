package main

import (
	"net/http"

	"github.com/aykahs/Gin-Boilerplate/configs"
	"github.com/aykahs/Gin-Boilerplate/pkg/mysql"
	"github.com/aykahs/Gin-Boilerplate/router"
	"github.com/gin-gonic/gin"

	uuid "github.com/google/uuid"
)

var db = make(map[string]string)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(RequestIDMiddleware())
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func main() {
	configs.Init()
	EnvConfig := configs.EnvConfig
	mysql.Connect(&EnvConfig.Mysql)
	router.Init()
	r := router.Router
	r.Run(EnvConfig.Server.Port)
}
