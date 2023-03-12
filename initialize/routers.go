package initialize

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/router"
)

// Routers The detailed information:
// @Title Routers
// @Description 路由配置信息
// @Return *gin.Engine
func Routers() *gin.Engine {
	Router := gin.Default()
	group := Router.Group("")
	router.RouterGroupApp.ShopRouter.InitShopRouter(group)
	router.RouterGroupApp.BlogRouter.InitBlogRouter(group)
	router.RouterGroupApp.VoucherRouter.InitVoucherRouter(group)
	router.RouterGroupApp.UserRouter.InitUserRouter(group)
	router.RouterGroupApp.UploadRouter.InitUploadRouter(group)
	router.RouterGroupApp.FollowRouter.InitFollowRouter(group)
	router.RouterGroupApp.CommentRouter.InitCommentRouter(group)
	return Router
}
