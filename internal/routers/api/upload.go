package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/internal/service"
	"github.com/lackone/gin-blog/pkg/app"
	"github.com/lackone/gin-blog/pkg/convert"
	"github.com/lackone/gin-blog/pkg/errcode"
	"github.com/lackone/gin-blog/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	res := app.NewResponse(c)
	fileHeader, err := c.FormFile("file")
	fileType := convert.Str(c.PostForm("type")).MustInt()
	if err != nil {
		errs := errcode.InvalidParams.WithDetails(err.Error())
		res.ToErrorResponse(errs)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		res.ToErrorResponse(errcode.InvalidParams)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		errs := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		res.ToErrorResponse(errs)
		return
	}
	newService := service.NewService(c)
	uploadFile, err := newService.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		errs := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		res.ToErrorResponse(errs)
		return
	}
	res.ToResponse(gin.H{
		"url": uploadFile.Url,
	})
}
