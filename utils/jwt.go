package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-gin-shop/enter/sys"
	"go-gin-shop/enter/tb"
	"go-gin-shop/global"
)

type JwtUtils struct {
}

func (JwtUtils) CreateToken(jwtClaims sys.JwtClaims) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims)
	token, err := claims.SignedString(global.JWTKey)
	return token, err
}

func (JwtUtils) ParseToken(token string) (*sys.JwtClaims, error) {
	jwtClaims := &sys.JwtClaims{}
	_, _, err := jwt.NewParser().ParseUnverified(token, jwtClaims)
	if err != nil {
		return nil, err
	}
	_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return global.JWTKey, nil
	})
	if err != nil {
		return nil, err
	}
	return jwtClaims, nil
}

func (JwtUtils) JwtGetUser(c *gin.Context) *tb.TbUser {
	// 获取jwt
	auth := c.GetHeader("authorization")
	// 从redis中获取值
	result, err := global.RedisDb.HGetAll(global.Content, auth).Result()
	if err != nil {
		return nil
	}
	// map转化为user结构体
	user := EnterUtilsApp.StructMapUtils.MapToStruct(&result)
	return user
}
