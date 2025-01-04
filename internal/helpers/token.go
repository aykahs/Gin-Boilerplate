package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func getToken(ctx *gin.Context) string {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		ctx.Abort()
		return ""
	}
	parts := strings.Split(tokenStr, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	token := parts[1]
	return token
}
