package controllers

import (
	"net/http"

	"github.com/aykahs/Gin-Boilerplate/internal/services/keyclock"
	"github.com/gin-gonic/gin"
)

var auth = new(keyclock.KeyclockAuthService)

type AuthController struct{}

func (authController *AuthController) Login(ctx *gin.Context) {

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

	token, err := auth.Login(username, password)
	if err != nil || token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Username or Password Error"})
		return
	}
	ctx.String(http.StatusOK, token)

}
