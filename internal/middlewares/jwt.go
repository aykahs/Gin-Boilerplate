package middlewares

import (
	"net/http"

	"github.com/aykahs/Gin-Boilerplate/internal/services/utils"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString, err := utils.GetToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
		}
		claims, err := utils.JwtKeyClockVerify(tokenString)
		ctx.Set("auth", claims)

		if err != nil || claims == nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		ctx.Next()
	}

}
