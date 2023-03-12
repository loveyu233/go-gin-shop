package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gin-shop/enter/res"
	"go-gin-shop/global"
	"go-gin-shop/utils"
	"time"
)

type CaptchaApi struct {
}

// 获取滑块验证码
func (CaptchaApi) GetCaptchaCode(c *gin.Context) {
	phone := c.Query("phone")

	if !utils.EnterUtilsApp.PhoneCodeUtils.CheckMobile(phone) {
		res.Err(c, "手机格式错误")
		return
	}
	captchaModel := utils.CreateCode()
	modelToVo := utils.CaptchaModelToVo(&captchaModel)
	global.RedisDb.Set(context.TODO(), phone, captchaModel.X, time.Second*120)
	res.OkData(c, modelToVo)
}

// 验证滑块验证码
func (CaptchaApi) VerificationCaptchaCode(c *gin.Context) {
	phone := c.Query("phone")
	if !utils.EnterUtilsApp.PhoneCodeUtils.CheckMobile(phone) {
		res.Err(c, "手机格式错误")
		return
	}
	xValue := c.Query("x")
	if global.RedisDb.Get(context.TODO(), phone).Val() != xValue {
		res.Err(c, "验证码错误")
		return
	}
	res.Ok(c)
}
