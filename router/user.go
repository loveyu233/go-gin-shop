package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/api"
	"go-gin-shop/middleware"
)

type UserRouter struct {
}

// InitUserRouter The detailed information:
// @Title InitUserRouter
// @Description user路由初始化
// @Param Group
func (UserRouter) InitUserRouter(Group *gin.RouterGroup) {
	user := Group.Group("/user")
	// 验证码
	user.POST("/code", api.ApiGroupApp.UserApi.Code)
	// 登陆
	user.POST("/login", api.ApiGroupApp.UserApi.Login)
	// 退出
	user.POST("/logout", api.ApiGroupApp.UserApi.LogOut)
	// 查询指定用户
	user.GET("/:userId", api.ApiGroupApp.UserApi.SelectById)
	authGroup := user.Group("")
	authGroup.Use(middleware.Auth)
	// 查询用户
	authGroup.GET("/info/:id", api.ApiGroupApp.UserApi.Info)
	// 查询登陆用户
	authGroup.GET("/me", api.ApiGroupApp.UserApi.Me)
	// 签到
	authGroup.GET("/sign", api.ApiGroupApp.UserApi.Sign)
	// 签到统计
	authGroup.GET("/signCount", api.ApiGroupApp.UserApi.SignCount)
}
