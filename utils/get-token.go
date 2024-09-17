package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractTokenFromHeader(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader("Authorization")
	parts := strings.Split(header, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}
