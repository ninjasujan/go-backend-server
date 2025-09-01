package middleware

import (
	"app/server/common/logger"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path

		// Log the request method and path
		// You can replace this with your preferred logging mechanism
		logger.Debug().
			Str("method", method).
			Str("path", path).
			Msg("Incoming request: method=" + method + " - path=" + path)

		// Proceed to the next middleware or handler
		c.Next()
	}
}
