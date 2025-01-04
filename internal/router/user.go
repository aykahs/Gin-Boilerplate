package router

import (
	"github.com/aykahs/Gin-Boilerplate/internal/controllers"
	"github.com/gin-gonic/gin"
)

var userController = new(controllers.UserController)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {

	user := r.Group("/users")
	{
		// user.POST("/login", userController.Login)
		// user.POST("/regiser", userController.Register)
		// user.GET("/list", userController.Register)
		// user.Use(middlewares.Jwt())
		// {
		// 	user.GET("/me", userController.AuthMe)
		// }

	}
	return user
}
