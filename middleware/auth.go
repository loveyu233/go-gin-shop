package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/enter/res"
	"go-gin-shop/utils"
)

func Auth(c *gin.Context) {
	// 获取jwt
	auth := c.GetHeader("authorization")
	jwtClaims, err := utils.EnterUtilsApp.JwtUtils.ParseToken(auth)
	if err != nil {
		res.Err(c, "jwt解析失败")
		c.Abort()
		return
	}
	c.Set("userid", jwtClaims.Uid)
}
