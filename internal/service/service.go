package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/internal/dao"
)

type Service struct {
	ctx *gin.Context
	dao *dao.Dao
}

func NewService(ctx *gin.Context) *Service {
	return &Service{
		ctx: ctx,
		dao: dao.NewDao(global.DB.WithContext(ctx.MustGet("SpanCtx").(context.Context))),
	}
}
