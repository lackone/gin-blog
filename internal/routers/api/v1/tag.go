package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/internal/service"
	"github.com/lackone/gin-blog/pkg/app"
	"github.com/lackone/gin-blog/pkg/convert"
	"github.com/lackone/gin-blog/pkg/errcode"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {

}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param status query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	req := service.TagListRequest{}
	res := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToErrorResponse(errs)
		return
	}
	newService := service.NewService(c)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	count, err := newService.CountTag(&service.CountTagRequest{Name: req.Name, Status: req.Status})
	if err != nil {
		res.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	list, err := newService.GetTagList(&req, &pager)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	res.ToResponseList(list, count)
	return
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param status body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	req := service.CreateTagRequest{}
	res := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToErrorResponse(errs)
		return
	}
	newService := service.NewService(c)
	err := newService.CreateTag(&req)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	res.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param status body int false "状态" Enums(0, 1) default(1)
// @Param updated_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	req := service.UpdateTagRequest{
		Id: convert.Str(c.Param("id")).MustUInt(),
	}
	res := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToErrorResponse(errs)
		return
	}
	newService := service.NewService(c)
	err := newService.UpdateTag(&req)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	res.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	req := service.DeleteTagRequest{
		Id: convert.Str(c.Param("id")).MustUInt(),
	}
	res := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToErrorResponse(errs)
		return
	}
	newService := service.NewService(c)
	err := newService.DeleteTag(&req)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	res.ToResponse(gin.H{})
	return
}
