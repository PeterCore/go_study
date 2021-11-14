package routers

import (
	"my-blog-sevice/global"
	"my-blog-sevice/internal/middleware"
	"my-blog-sevice/internal/routers/api"
	v1 "my-blog-sevice/internal/routers/api/v1"
	"net/http"

	_ "my-blog-sevice/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UpLoadFiles)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv := r.Group("/api/v1")
	{
		apiv.POST("/tags", tag.Create)
		apiv.DELETE("/tags/:id", tag.Delete)
		apiv.PUT("/tags/:id", tag.Update)
		apiv.PATCH("/tags/:id/state", tag.Update)
		apiv.GET("/tags", tag.List)

		apiv.POST("/articles", article.Create)
		apiv.DELETE("/articles/:id", article.Delete)
		apiv.PUT("/articles/:id", article.Update)
		apiv.PATCH("/articles/:id/state", article.Update)
		apiv.GET("/articles", article.List)

	}
	return r

}
