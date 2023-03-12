package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-gin-shop/global"
	"time"
)

type RedisBitmapUtil struct {
}

func (RedisBitmapUtil) Sign(userId int) bool {
	key := fmt.Sprintf("%s%d:%d-%s", global.RedisSign, userId, time.Now().Year(), time.Now().Format("01"))
	day := time.Now().Day()
	result, err := global.RedisDb.SetBit(global.Content, key, int64(day), 1).Result()
	if err != nil || result <= 0 {
		return false
	}
	return true
}

func (RedisBitmapUtil) SignCount(userId int, year, moth string) int64 {
	key := fmt.Sprintf("%s%d:%s-%s", global.RedisSign, userId, year, moth)
	result, err := global.RedisDb.BitCount(global.Content, key, &redis.BitCount{
		Start: 0,
		End:   -1,
	}).Result()
	if err != nil {
		return 0
	}
	return result
}
