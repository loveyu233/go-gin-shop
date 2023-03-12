package initialize

import "github.com/go-redis/redis/v8"

// InitRedis The detailed information:
// @Title InitRedis
// @Description redis连接信息
// @Param dsn 连接配置
// @Return *redis.Client
func InitRedis(dsn string) *redis.Client {
	url, err := redis.ParseURL(dsn)
	if err != nil {
		panic("redis初始化错误：" + err.Error())
		return nil
	}
	client := redis.NewClient(url)
	return client
}
