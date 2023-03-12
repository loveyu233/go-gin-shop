package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/api"
	"go-gin-shop/middleware"
)

// BlogRouter blog路由集中注册
type BlogRouter struct{}

// InitBlogRouter The detailed information:
// @Title InitBlogRouter
// @Description 初始化blog路由
// @Param Router
func (BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	blog := Router.Group("blog")
	BlogApi := api.BlogApi{}
	// 发布文章
	blog.POST("", BlogApi.Blog)
	// blog点赞排行榜
	blog.GET("/hot", BlogApi.Hot)
	// 查看其他用户的blog
	blog.GET("/of/user", BlogApi.OfUser)
	// 点赞
	blog.GET("/likes/:id", BlogApi.LikesBlogId)
	// 获取指定id的blog
	blog.GET("/:id", BlogApi.ById)
	blogUser := blog.Group("", middleware.Auth)
	blogUser.PUT("/like/:id", BlogApi.LikeBlog)
	// 查询登陆账号的blog
	blogUser.GET("/of/me", BlogApi.OfMe)
	//用户关注消息
	blogUser.GET("/of/follow", BlogApi.OfFollow)
}
