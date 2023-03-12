package redisLock

import (
	"github.com/go-redis/redis/v8"
	"go-gin-shop/global"
	"log"
	"time"
)

const lockPrefix = "lock:"

// RedisLock redis 分布式锁
type RedisLock struct {
	LockName  string
	LockValue string
	TTl       time.Duration
}

func NewRedisLock(lockName, lockValue string, tll time.Duration) *RedisLock {
	return &RedisLock{
		LockName:  lockName,
		LockValue: lockValue,
		TTl:       tll,
	}
}

func (r *RedisLock) TryLock() bool {
	result, err := global.RedisDb.SetNX(global.Content, lockPrefix+r.LockName, r.LockValue, r.TTl).Result()
	if err != nil {
		log.Println(err)
	}
	return result
}

func (r *RedisLock) UnLock() {
	script := redis.NewScript(global.UnLockLuaScript)
	script.Run(global.Content, global.RedisDb, []string{lockPrefix + r.LockName}, r.LockValue)
}
