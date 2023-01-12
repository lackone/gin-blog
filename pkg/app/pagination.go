package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.Str(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.Str(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		pageSize = global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		pageSize = global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetOffset(page, pageSize int) int {
	offset := 0
	if page > 0 {
		offset = (page - 1) * pageSize
	}
	return offset
}
