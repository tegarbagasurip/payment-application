package middleware

import (
	"net/http"
	"strings"

	"payment-application/utils/security"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header authHeader
		err := c.ShouldBindHeader(&header)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid header " + err.Error(),
			})
			c.Abort()
			return
		}

		tokenString := strings.Replace(header.AuthorizationHeader, "Bearer ", "", 1)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token string",
			})
			c.Abort()
			return
		}

		claims, err := security.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
