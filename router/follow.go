package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/api"
)

type FollowRouter struct {
}

func (FollowRouter) InitFollowRouter(router *gin.RouterGroup) {
	follow := router.Group("/follow")
	// 关注 vlogID：关注的user，isFollow：是否关注
	follow.PUT("/:blogUserId/:isFollow", api.ApiGroupApp.FollowApi.Follow)
	// 是否关注blogId
	follow.GET("/or/not/:blogUserId", api.ApiGroupApp.FollowApi.FollowOrNot)
	// 共同关注
	follow.GET("/common/:blogId", api.ApiGroupApp.FollowApi.Common)
}
