package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lackone/gin-blog/internal/service"
	"github.com/lackone/gin-blog/pkg/app"
	"github.com/lackone/gin-blog/pkg/convert"
	"github.com/lackone/gin-blog/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取单个文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	req := service.GetArticleRequest{
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
	article, err := newService.GetArticle(&req)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}
	res.ToResponse(article)
	return
}

// @Summary 获取多个文章
// @Produce json
// @Param name query string false "文章名称"
// @Param tag_id query int false "标签ID"
// @Param status query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	req := service.ArticleListRequest{}
	res := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToErrorResponse(errs)
		return
	}
	newService := service.NewService(c)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	list, count, err := newService.GetArticleList(&req, &pager)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}
	res.ToResponseList(list, count)
	return
}

// @Summary 创建文章
// @Produce json
// @Param tag_id body string true "标签ID"
// @Param title body string true "文章标题"
// @Param desc body string false "文章简述"
// @Param cover body string true "封面图片"
// @Param content body string true "文章内容"
// @Param created_by body int true "创建者"
// @Param status body int false "状态"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	req := service.CreateArticleRequest{}
	res := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &req)
	if !valid {
		errs := errcode.InvalidParams.WithDetails(errors.Errors()...)
		res.ToErrorResponse(errs)
		return
	}
	newService := service.NewService(c)
	err := newService.CreateArticle(&req)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	res.ToResponse(gin.H{})
	return
}

// @Summary 更新文章
// @Produce json
// @Param tag_id body string false "标签ID"
// @Param title body string false "文章标题"
// @Param desc body string false "文章简述"
// @Param cover body string false "封面图片"
// @Param content body string false "文章内容"
// @Param updated_by body string true "修改者"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	req := service.UpdateArticleRequest{
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
	err := newService.UpdateArticle(&req)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	res.ToResponse(gin.H{})
	return
}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	req := service.DeleteArticleRequest{
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
	err := newService.DeleteArticle(&req)
	if err != nil {
		res.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	res.ToResponse(gin.H{})
	return
}
