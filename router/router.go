package router

import (
	"github.com/aykahs/Gin-Boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	// Global middlewares
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())
	LoadUserRoutes(Router)

}
