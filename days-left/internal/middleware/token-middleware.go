package middleware

import "github.com/gin-gonic/gin"

func TokenMiddlerware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		initToken := "Bearer 2165456654651651651"

		if token != initToken {
			c.AbortWithStatusJSON(401, map[string]interface{}{
				"error": "unauthorized",
			})
		}

		c.Next()
	}
}
