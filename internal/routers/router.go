package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lackone/gin-blog/docs" //千万别忘了导入生成的docs
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/internal/middleware"
	"github.com/lackone/gin-blog/internal/routers/api"
	v1 "github.com/lackone/gin-blog/internal/routers/api/v1"
	"github.com/lackone/gin-blog/pkg/limit"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimit = limit.NewMethodLimit().AddBuckets(limit.LimitBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		//访问日志
		r.Use(middleware.AccessLog())
		//自定义Recovery
		r.Use(middleware.Recovery())
	}

	//链路追踪
	r.Use(middleware.Trace(global.ServerSetting.ServiceName))
	//限流中间件
	r.Use(middleware.Limit(methodLimit))
	//超时中间件
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	//翻译中间件
	r.Use(middleware.Trans())
	//添加公共信息
	r.Use(middleware.AppInfo())

	//swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()

	//文件上传
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/uploads", http.Dir(global.AppSetting.UploadSavePath))

	//获取token
	r.GET("/auth", api.GetAuth)

	group := r.Group("/api/v1")
	group.Use(middleware.JWT())
	{
		//增加标签
		group.POST("/tags", tag.Create)
		//删除标签
		group.DELETE("/tags/:id", tag.Delete)
		//修改标签
		group.PUT("/tags/:id", tag.Update)
		group.PATCH("/tags/:id/state", tag.Update)
		//标签列表
		group.GET("/tags", tag.List)

		//添加文章
		group.POST("/articles", article.Create)
		//删除文章
		group.DELETE("/articles/:id", article.Delete)
		//修改文章
		group.PUT("/articles/:id", article.Update)
		group.PATCH("/articles/:id/state", article.Update)
		//文章列表
		group.GET("/articles", article.List)
		//文章详情
		group.GET("/articles/:id", article.Get)
	}

	return r
}
