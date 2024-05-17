package router

import (
	"net/http"

	"github.com/aykahs/Gin-Boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	LoadUserRoutes(Router)

}
