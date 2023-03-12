package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/enter/dto"
	"go-gin-shop/global"
	"go-gin-shop/service"
	"go-gin-shop/utils"
	"strconv"
)

type UserApi struct {
}

var userService = service.EnterServicesApp.UserService

// Info The detailed information:
// @Title Info
// @Description 查询用户相关信息
// @Param c
func (UserApi) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, userService.Info(id))
}

// Code The detailed information:
// @Title Code
// @Description 注册账号
// @Param c
func (UserApi) Code(c *gin.Context) {
	phone := c.Query("phone")
	c.JSON(200, userService.Code(phone, c))
}

// Login The detailed information:
// @Title Login
// @Description 用户登录
// @Param c 手机号+验证码 / 手机号+密码
func (UserApi) Login(c *gin.Context) {
	loginUser := dto.LoginUser{}
	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(200, dto.Err("参数解析错误"))
		return
	}
	c.JSON(200, userService.Login(loginUser, c))
}

// LogOut The detailed information:
// @Title LogOut
// @Description 账号退出
// @Param c
func (UserApi) LogOut(c *gin.Context) {
	auth := c.GetHeader("authorization")
	global.RedisDb.Del(global.Content, auth)
	c.JSON(200, dto.Ok())
}

// Me The detailed information:
// @Title Me
// @Description 获取当前登录的用户并返回
// @Param c
func (UserApi) Me(c *gin.Context) {
	c.JSON(200, userService.Me(c))
}

func (UserApi) SelectById(c *gin.Context) {
	userId := c.Param("userId")
	c.JSON(200, userService.ById(userId))
}

func (UserApi) Sign(c *gin.Context) {
	c.JSON(200, userService.Sign(int(utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID)))
}

func (UserApi) SignCount(c *gin.Context) {
	year := c.Query("year")
	month := c.Query("month")
	c.JSON(200, userService.SignCount(int(utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID), year, month))
}
