package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/tb"
	"go-gin-shop/service"
	"go-gin-shop/utils"
	"strconv"
)

type BlogApi struct{}

// 获取点赞最多的
func (BlogApi) Hot(c *gin.Context) {
	current := c.Query("current")
	currentInt, _ := strconv.Atoi(current)
	userId := utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID
	c.JSON(200, service.EnterServicesApp.BlogService.Hot(currentInt, userId))
}

// LikeBlog The detailed information:
// @Title likeBlog
// @Description 点赞➕1
// @Param c
func (BlogApi) LikeBlog(c *gin.Context) {
	blogId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, dto.Err("blog id err"))
		return
	}
	userId := utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID
	c.JSON(200, service.EnterServicesApp.BlogService.LikeBlog(blogId, int(userId)))
}

// ById The detailed information:
// @Title ById
// @Description 返回具体blog信息
// @Param c
func (BlogApi) ById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, service.EnterServicesApp.BlogService.ById(id))
}

func (BlogApi) OfMe(c *gin.Context) {
	c.JSON(200, service.EnterServicesApp.BlogService.OfMe(c))
}

// 上传文章
func (BlogApi) Blog(c *gin.Context) {
	blog := tb.TbBlog{}
	if c.ShouldBindJSON(&blog) != nil {
		c.JSON(200, dto.Err("格式错误"))
		return
	}
	user := utils.EnterUtilsApp.JwtUtils.JwtGetUser(c)
	c.JSON(200, service.EnterServicesApp.BlogService.Blog(&blog, user))
}

// 点赞排行榜,这里的id是blog的id
func (BlogApi) LikesBlogId(c *gin.Context) {
	blogId := c.Param("id")
	c.JSON(200, service.EnterServicesApp.BlogService.BlogLikes(blogId))
}

// 获取用户id查blog
func (BlogApi) OfUser(c *gin.Context) {
	userIdString := c.Query("id")
	current := c.Query("current")
	c.JSON(200, service.EnterServicesApp.BlogService.OfUser(userIdString, current))
}

// 滑动获取blog
func (BlogApi) OfFollow(c *gin.Context) {
	lastId := c.Query("lastId")
	offset := c.DefaultQuery("offset", "0")
	c.JSON(200, service.EnterServicesApp.BlogService.OfFollowService(lastId, offset, utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID))
}
