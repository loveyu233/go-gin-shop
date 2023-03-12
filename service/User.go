package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/sys"
	"go-gin-shop/enter/tb"
	"go-gin-shop/global"
	"go-gin-shop/utils"

	"strconv"
	"time"
)

// User 用户信息
type User struct {
}

// UserById The detailed information:
// @Title UserById
// @Description 根据用户id进行查询
// @Param id
// @Return tb.TbUser
func (User) UserById(id int) tb.TbUser {
	user := tb.TbUser{}
	global.MysqlDb.Model(&user).Scopes(EnterServicesApp.PaginateService.byId(id)).Find(&user)
	return user
}

// Info The detailed information:
// @Title Info
// @Description 查询用户相关信息
// @Param id
// @Return tb.TbUser
func (User) Info(id int) dto.Response {
	if id == 0 {
		return dto.Err("请登录")
	}
	user := tb.TbUser{}
	global.MysqlDb.Model(&user).Scopes(EnterServicesApp.PaginateService.byId(id)).Find(&user)
	userVo := dto.UserModelToUserVo(user)
	return dto.OkData(userVo)
}

// Code The detailed information:
// @Title Code
// @Description 通过手机号注册账号
// @Param phone
func (User) Code(phone string, c *gin.Context) dto.Response {
	if !utils.EnterUtilsApp.PhoneCodeUtils.CheckMobile(phone) {
		return dto.Err("手机号格式不正确")
	}
	code := utils.EnterUtilsApp.PhoneCodeUtils.Code()
	// 把验证码存储到redis中
	global.RedisDb.Set(global.Content, phone, code, global.PhoneCodeTTLTime)
	return dto.Ok()
}

// Login The detailed information:
// @Title Login
// @Description 登陆
// @Param loginUser
// @Param c
// @Return dto.Response
func (User) Login(loginUser dto.LoginUser, c *gin.Context) dto.Response {
	// 验证手机号码格式是否正确
	if !utils.EnterUtilsApp.PhoneCodeUtils.CheckMobile(loginUser.Phone) {
		return dto.Err("手机号码格式不正确")
	}
	// 从redis中读取手机号验证码
	phoneCode, err := global.RedisDb.Get(global.Content, loginUser.Phone).Result()
	if err != nil {
		return dto.Err("code nil")
	}
	// 判断验证码是否正确
	if phoneCode == "" || phoneCode != loginUser.Code {
		return dto.Err("验证码错误")
	}
	user := &tb.TbUser{}
	// 查询该手机是否存在于数据库，不存在说明该用户为新用户
	global.MysqlDb.Model(user).Where("phone = ?", loginUser.Phone).Find(user)
	// 名字等于空说明该用户为新用户，则为其设置基本属性
	if user.NickName == "" {
		user.Phone = loginUser.Phone
		user.NickName = global.UserNiceNamePrefix + loginUser.Phone
		user.CreateTime = time.Now()
		// 存储该用户
		global.MysqlDb.Model(user).Omit("update_time").Create(&user)
	}
	// 因为要把数据发送到前端为了安全所以要把手机号和密码设置为空
	user.Phone = ""
	user.Password = ""
	// 创建jwt作为存储redis的key
	token, err := utils.EnterUtilsApp.JwtUtils.CreateToken(sys.JwtClaims{
		Uid: uint(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(global.TokenTTlTIme))),
		},
	})
	if err != nil {
		return dto.Err("jwt err")
	}
	// 用户数据存储到redis中
	global.RedisDb.HSet(global.Content, token, utils.EnterUtilsApp.StructMapUtils.StructToMap(*user))
	// 设置过期时间
	global.RedisDb.Expire(global.Content, token, global.RedisTokenTTLTIme)
	// 登陆成功删除redis中验证码
	global.RedisDb.Del(global.Content, loginUser.Phone)
	// 返回jwt给前端，前端下次访问则拿着jwt来获取信息
	return dto.OkData(token)
}

// Me The detailed information:
// @Title Me
// @Description 获取用户详细信息
// @Param c
// @Return dto.Response
func (User) Me(c *gin.Context) dto.Response {
	// 获取jwt
	auth := c.GetHeader("authorization")
	// 用jwt作为key获取redis中的value
	result, err := global.RedisDb.HGetAll(global.Content, auth).Result()
	if err != nil {
		return dto.Err("未登陆请登录")
	}
	// 返回数据
	return dto.OkData(result)
}

func (User) ById(userId string) dto.Response {
	user := tb.TbUser{}
	userIdInt, _ := strconv.Atoi(userId)
	affected := global.MysqlDb.Scopes(EnterServicesApp.PaginateService.byId(userIdInt)).First(&user).RowsAffected
	if affected <= 0 {
		return dto.Err("没有该用户")
	}
	return dto.OkData(user)
}

func (User) Sign(userId int) dto.Response {
	utils.EnterUtilsApp.RedisBitmapUtil.Sign(userId)
	return dto.Ok()
}

func (User) SignCount(userId int, year, moth string) dto.Response {
	return dto.OkData(utils.EnterUtilsApp.RedisBitmapUtil.SignCount(userId, year, moth))
}
