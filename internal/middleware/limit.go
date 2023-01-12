package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/pkg/app"
	"github.com/lackone/gin-blog/pkg/errcode"
	"github.com/lackone/gin-blog/pkg/limit"
)

// 限流
func Limit(l limit.LimitIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				res := app.NewResponse(c)
				res.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
