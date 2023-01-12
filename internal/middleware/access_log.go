package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (this *AccessLogWriter) Write(b []byte) (int, error) {
	//保存一份数据到我们自已的结构体里
	this.body.Write(b)
	return this.ResponseWriter.Write(b)
}

// 访问日志
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := &AccessLogWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = writer

		start := time.Now().Format("2006-01-02 15:04:05")
		c.Next()
		end := time.Now().Format("2006-01-02 15:04:05")

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"url":      c.Request.URL.String(),
			"response": writer.body.String(),
		}

		global.Logger.WithFields(fields).Infof(c, "access log: method %s status_code %d start %s end %s", c.Request.Method, writer.Status(), start, end)
	}
}
