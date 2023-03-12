package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/service"
	"go-gin-shop/utils"

	"strconv"
)

type FollowApi struct {
}

func (FollowApi) Follow(c *gin.Context) {
	blogUserId := c.Param("blogUserId")
	isFollowString := c.Param("isFollow")
	var isFollowBool bool
	isFollowBool, _ = strconv.ParseBool(isFollowString)
	userId := utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID
	c.JSON(200, service.EnterServicesApp.FollowService.Follow(blogUserId, isFollowBool, int(userId)))
}

func (FollowApi) FollowOrNot(c *gin.Context) {
	blogUserId := c.Param("blogUserId")
	userId := utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID
	c.JSON(200, service.EnterServicesApp.FollowService.FollowOrNot(blogUserId, int(userId)))
}

func (FollowApi) Common(c *gin.Context) {
	blogId := c.Param("blogId")
	c.JSON(200, service.EnterServicesApp.FollowService.Common(blogId, int(utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID)))
}
