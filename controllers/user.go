package controllers

import (
	"net/http"

	"github.com/aykahs/Gin-Boilerplate/services"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

type UserController struct{}

func (userController *UserController) AuthMe(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	var userIDInterface interface{}
	userIDInterface = userID
	user_ID, ok := userIDInterface.(uint)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "something went wrong!"})
	}
	err, user := userService.Me(user_ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}
func (userController *UserController) Register(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	user, err := userService.Register(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (userController *UserController) Login(ctx *gin.Context) {

	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	username := data["username"].(string)
	password := data["password"].(string)

	if username == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Username and Password is required"})
		return
	}

	token, err := userService.LoginByUsernamePassword(username, password)
	if err != nil || token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Username or Password Error"})
		return
	}

	ctx.String(http.StatusOK, token)

}
