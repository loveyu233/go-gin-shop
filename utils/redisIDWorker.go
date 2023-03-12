package utils

import (
	"fmt"
	"go-gin-shop/global"
	"time"
)

type RedisIDWorker struct {
}

// NextId The detailed information:
// @Title NextId
// @Description 生成全局唯一ID
// @Param keyPrefix 前缀
// @Return int64 64位是应为：第一位为符号位，后边31位为时间戳，最后32位为序列号（也就是redis自增的到的）
func (RedisIDWorker) NextId(keyPrefix string) int64 {
	// 生成时间戳
	timestamp := time.Now().Unix()
	// 获取当前的年月日
	timeData := time.Now().Format("2006:01:02")
	// redis自增，key拼接为：icr + key前缀 + 当前年月日
	count, _ := global.RedisDb.Incr(global.Content, fmt.Sprintf("%s:%s:%s:", "icr", keyPrefix, timeData)).Result()
	// 左移32位，为序列号留出位置，然后 或运算 上序列号就得到最后的全局唯一ID
	return timestamp<<32 | count
}
