package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/api"
)

func InitCaptchaRouter(group *gin.RouterGroup) {
	captchaApi := api.ApiGroupApp.CaptchaApi
	captcha := group.Group("/captcha")
	captcha.GET("/get", captchaApi.GetCaptchaCode)
	captcha.POST("/check", captchaApi.VerificationCaptchaCode)
}
