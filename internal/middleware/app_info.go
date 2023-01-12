package middleware

import "github.com/gin-gonic/gin"

// app公共信息
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app-name", "gin-blog")
		c.Set("app-version", "1.0")
		c.Next()
	}
}
