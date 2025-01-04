package middlewares

import (
	"net/http"
	"strings"

	"github.com/aykahs/Gin-Boilerplate/internal/services/utils"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		parts := strings.Split(tokenStr, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}

		token := parts[1]
		claims, err := utils.JwtVerify(token)
		// fmt.Println("here")
		// fmt.Println(claims.Uid)
		ctx.Set("user_id", claims.Uid)

		if err != nil || claims == nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		ctx.Next()
	}

}
