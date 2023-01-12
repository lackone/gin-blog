package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/internal/service"
	"github.com/lackone/gin-blog/pkg/app"
	"github.com/lackone/gin-blog/pkg/errcode"
	"go.opentelemetry.io/otel/attribute"
)

func GetAuth(c *gin.Context) {
	p := service.AuthRequest{}
	res := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &p)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToErrorResponse(errs)
		return
	}
	newService := service.NewService(c)
	err := newService.CheckAuth(&p)
	if err != nil {
		res.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	jwt, err := app.MakeJwt(p.AppKey, p.AppSecret)
	if err != nil {
		res.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	TestJaeger(c)

	res.ToResponse(gin.H{
		"token": jwt,
	})
}

func TestJaeger(c *gin.Context) {
	_, span := global.Trace.Start(c.MustGet("SpanCtx").(context.Context), "test")
	span.SetAttributes(attribute.Key("TestJaeger").String("TestJaeger"))
	defer span.End()
}
