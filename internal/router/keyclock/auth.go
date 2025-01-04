package router

import (
	"github.com/aykahs/Gin-Boilerplate/internal/controllers/keyclock"
	"github.com/gin-gonic/gin"
)

var authController = new(keyclock.AuthController)

func LoadKeyclockRoutes(r *gin.Engine) *gin.RouterGroup {

	Keyclock := r.Group("/Keyclocks")
	{
		Keyclock.POST("/login", authController.Login)
		// Keyclock.POST("/regiser", authController.Register)
		// Keyclock.GET("/list", authController.Register)
		// Keyclock.Use(middlewares.Jwt())
		// {
		// 	Keyclock.GET("/me", authController.AuthMe)
		// }

	}
	return Keyclock
}
