package authcontrollers

import (
	"net/http"

	keyclockservice "github.com/aykahs/Gin-Boilerplate/internal/services/keyclock"
	"github.com/aykahs/Gin-Boilerplate/internal/services/utils"
	"github.com/gin-gonic/gin"
)

var auth = &keyclockservice.KeyClockAuthService{
	HttpCurl: &keyclockservice.HttpCurl{},
}

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

func (authController *AuthController) KeyClockRefresh(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	refresh_token := data["refresh_token"].(string)
	tokenResponse, err := auth.Refresh(refresh_token)

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
	token, err := utils.GetToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
	tokenResponse, err := auth.Me(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user_info": tokenResponse,
	})
}
