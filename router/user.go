package router

import (
	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {

	user := r.Group("/users")
	{
		// user.POST("/loginByUsernamePassword", userController.LoginByUsernamePassword)
	}
	return user
}
