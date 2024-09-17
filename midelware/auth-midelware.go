package middleware

import (
	"backend-jona-golang/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := utils.ExtractTokenFromHeader(ctx)

		if err != nil || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Token is missing or invalid",
			})

			ctx.Abort()
			return
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid token",
			})

			ctx.Abort()
			return
		}

		userID := fmt.Sprintf("%d", claims.ID)
		userRole := claims.Role
		userEmail := claims.Email

		ctx.Set("userID", userID)
		ctx.Set("userRole", userRole)
		ctx.Set("userEmail", userEmail)

		ctx.Next()
	}
}
