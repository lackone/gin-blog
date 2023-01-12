package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// 超时控制
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeout, cancelFunc := context.WithTimeout(c.Request.Context(), t)
		defer cancelFunc()

		c.Request = c.Request.WithContext(timeout)
		c.Next()
	}
}
