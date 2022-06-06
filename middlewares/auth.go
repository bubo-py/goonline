package middlewares

import (
	"github.com/bubo-py/goonline/auth"
	"github.com/gin-gonic/gin"
)

// Auth is a middleware function to provide the secured routes
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		if err := auth.ValidateToken(tokenString); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
