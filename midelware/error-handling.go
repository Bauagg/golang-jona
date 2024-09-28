package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandlingMiddleware adalah middleware untuk menangani error secara global
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Menangani error secara global
		defer func() {
			if err := recover(); err != nil {
				// Log error
				log.Printf("Recovered from panic: %v", err)

				// Kirim respons error
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   true,
					"message": "Internal server error ===>",
				})
				c.Abort()
			}
		}()

		// Melanjutkan ke handler berikutnya
		c.Next()

		// Menangani error yang dikembalikan dari handler
		if len(c.Errors) > 0 {
			// Log error
			for _, e := range c.Errors {
				log.Printf("Error: %v", e.Err)
			}

			// Kirim respons error
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   true,
				"message": c.Errors.ByType(gin.ErrorTypePrivate).String(),
			})
			c.Abort()
		}
	}
}
