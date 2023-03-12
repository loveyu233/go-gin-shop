package utils

import (
	"github.com/go-redis/redis/v8"
	"go-gin-shop/enter/tb"
	"go-gin-shop/global"
	"gorm.io/gorm/clause"
	"time"
)

type RedisBlogLiked struct {
}

// IsExist The detailed information:
// @Title IsExist
// @Description 查询该blog的set集合中是否存在该用户id
// @Param key
// @Param value
// @Return bool
func (RedisBlogLiked) IsExist(key, value string) bool {
	f, err := global.RedisDb.ZScore(global.Content, key, value).Result()
	if err != nil {
		return false
	}
	if f > 0 {
		return true
	}
	return false
}

// LikedAdd The detailed information:
// @Title LikedAdd
// @Description 该blog的set中添加用户id
// @Param key
// @Param value
// @Return bool
func (RedisBlogLiked) LikedAdd(key string, value any) bool {
	i, err := global.RedisDb.ZAdd(global.Content, key, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: value,
	}).Result()
	if err != nil || i == 0 {
		return false
	}
	return true
}

// LikedDel The detailed information:
// @Title LikedDel
// @Description  删除blog对应set集合中用户id
// @Param key
// @Param value
// @Return bool
func (RedisBlogLiked) LikedDel(key string, value any) bool {
	i, err := global.RedisDb.ZRem(global.Content, key, value).Result()
	if err != nil || i == 0 {
		return false
	}
	return true
}

func (RedisBlogLiked) GetBlogLikedUser(key string) []*tb.TbUser {
	usersId, err := global.RedisDb.ZRange(global.Content, key, 0, 4).Result()
	if err != nil {
		return nil
	}
	users := make([]*tb.TbUser, 0)
	// 使用orderBy进行排序
	global.MysqlDb.Model(&tb.TbUser{}).Where("id IN ?", usersId).Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "FIELD(id,?)",
			Vars:               []interface{}{usersId},
			WithoutParentheses: true,
		},
	}).Find(&users)
	for i := range users {
		users[i].Password = ""
		users[i].Phone = ""
	}
	return users
}
