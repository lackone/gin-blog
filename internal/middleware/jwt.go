package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lackone/gin-blog/pkg/app"
	"github.com/lackone/gin-blog/pkg/errcode"
)

// jwt验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := app.NewResponse(c)
		token := ""
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			res.ToErrorResponse(errcode.InvalidParams)
			c.Abort()
			return
		} else {
			_, err := app.ParseJwt(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					res.ToErrorResponse(errcode.UnauthorizedTokenTimeout)
				default:
					res.ToErrorResponse(errcode.UnauthorizedTokenError)
				}
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
