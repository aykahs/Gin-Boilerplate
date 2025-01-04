package middlewares

import (
	"net/http"

	"github.com/aykahs/Gin-Boilerplate/internal/services/utils"
	"github.com/gin-gonic/gin"
	"github.com/aykahs/Gin-Boilerplate/internal/helpers"

)

func Jwt() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := helpers.getToken(ctx)
		claims, err := utils.JwtKeyClockVerify(token)
		ctx.Set("auth", claims)

		if err != nil || claims == nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		ctx.Next()
	}

}
