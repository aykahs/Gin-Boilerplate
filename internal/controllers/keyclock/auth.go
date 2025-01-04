package authcontrollers

import (
	"fmt"
	"net/http"

	keyclockservice "github.com/aykahs/Gin-Boilerplate/internal/services/keyclock"
	"github.com/gin-gonic/gin"
)

var auth = new(keyclockservice.KeyClockAuthService)

type AuthController struct{}

func (authController *AuthController) KeyClockLogin(ctx *gin.Context) {

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

	tokenResponse, err := auth.Login(username, password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokenResponse.AccessToken,
		"expires_in":    tokenResponse.ExpiresIn,
		"token_type":    tokenResponse.TokenType,
		"refresh_token": tokenResponse.RefreshToken,
	})
}

func (authController *AuthController) Me(ctx *gin.Context) {

	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	token := data["token"].(string)
	fmt.Println(token)
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "token expired"})
		return
	}

	tokenResponse, err := auth.Me(token)
	fmt.Println(err, "asjb")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user_info": tokenResponse,
	})
}
