package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/api"
	"go-gin-shop/middleware"
)

type UploadRouter struct {
}

func (UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {
	upload := Router.Group("/upload")
	upload.Use(middleware.Auth)
	{
		// 上传blog图片
		upload.POST("/blog", api.ApiGroupApp.UploadApi.Save)
		// 删除blog图片
		upload.GET("/delete", api.ApiGroupApp.UploadApi.Delete)
	}
}
