package authrouter

import (
	authcontrollers "github.com/aykahs/Gin-Boilerplate/internal/controllers/keyclock"
	"github.com/aykahs/Gin-Boilerplate/internal/middlewares"
	"github.com/gin-gonic/gin"
)

var authController = new(authcontrollers.AuthController)

func LoadKeyclockRoutes(r *gin.Engine) *gin.RouterGroup {

	Keyclock := r.Group("/")
	{
		Keyclock.POST("/login", authController.KeyClockLogin)
		Keyclock.Use(middlewares.Jwt())
		{
			Keyclock.POST("/refresh", authController.KeyClockRefresh)
			Keyclock.POST("/me", authController.Me)

		}

	}
	return Keyclock
}
