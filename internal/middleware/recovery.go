package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/pkg/app"
	"github.com/lackone/gin-blog/pkg/email"
	"github.com/lackone/gin-blog/pkg/errcode"
	"time"
)

// 捕获异常
func Recovery() gin.HandlerFunc {
	defEmail := email.NewDefaultEmail()
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCaller(2).WithCallersFrames().Errorf(c, "panic recover err %v", err)

				emailErr := defEmail.SendEmail(global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间 %s", time.Now().Format("2006-01-02 15:04:05")),
					fmt.Sprintf("错误信息 %v", err),
				)
				if emailErr != nil {
					global.Logger.Panicf(c, "mail.SendEmail err %v", emailErr)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
				return
			}
		}()

		c.Next()
	}
}
