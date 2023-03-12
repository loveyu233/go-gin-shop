package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/linvon/cuckoo-filter"
	"go-gin-shop/enter/config"
	"go-gin-shop/enter/sys"
	"gorm.io/gorm"
	"time"
)

const (
	// ConfigFilePath 配置文件位置
	ConfigFilePath = "config.yaml"

	// NaxPageSize 页面最大数量
	NaxPageSize = 10

	// BlogFileSaveLocal blog文件保存位置
	BlogFileSaveLocal = "/opt/homebrew/var/www/html/hmdp/imgs/blogs/"

	// UserNiceNamePrefix 用户名前缀
	UserNiceNamePrefix = "user_"

	// ShopCache 店铺信息缓存
	ShopCache = "cache:shop:"

	// ShopLock shop锁的key
	ShopLock = "lock:shop:"

	// ShopTypeListCache 商铺分类信息缓存
	ShopTypeListCache = "cache:shopType-list"

	// VoucherCache 优惠卷信息缓存
	VoucherCache = "cache:voucher:"

	// VoucherRedisKey voucher stock redisKey
	VoucherRedisKey = "seckill:stock:"

	// RedisCacheLiked 点赞set集合key前缀
	RedisCacheLiked = "blog:liked:"

	// RedisFeedFans 关注推送
	RedisFeedFans = "feed:"

	RedisShopGeo = "shop:geo:"

	RedisSign = "sign:"

	// TTLTime 缓存ttl时间
	TTLTime = time.Second * 60 * 30

	// PhoneCodeTTLTime 验证码过期时间
	PhoneCodeTTLTime = time.Second * 60

	// RedisTokenTTLTIme redis的token过期时间
	RedisTokenTTLTIme = time.Second * 60 * 29

	// RedisTryLockTTLTime redis实现互斥锁的ttl时间
	RedisTryLockTTLTime = time.Second * 10
)

var (
	// Config 全部配置
	Config = &config.Config{}
	// MysqlDb mysql连接
	MysqlDb *gorm.DB
	// RedisDb redis连接
	RedisDb *redis.Client
	// mongodb连接
	MongoDb *sys.MongoClient
	// Content content
	Content = context.TODO()
	// Bloomfilter 布隆过滤器
	Bloomfilter *cuckoo.Filter
	// JWTKey JWT密钥
	JWTKey = []byte("qaqaq122333..")

	// TokenTTlTIme JWT Token过期时间
	TokenTTlTIme = 6

	// TimeNow 当前时间
	TimeNow = time.Now().Format("2006-01-02 15:04:05.000")

	VoucherOrderRabbitMQQueueName = "voucherOrderCreate"
)

var (
	// UnLockLuaScript 释放锁的lua脚本
	UnLockLuaScript = `
	if redis.call("GET",KEYS[1]) == ARGV[1] then
	return redis.call("DEL",KEYS[1])
	else
		return 0
	end`

	// IsQualificationLuaScript 判断是否有抢购优惠卷资格的脚本
	IsQualificationLuaScript = `
	local voucherId = ARGV[1]
	local userId = ARGV[2]
	local stockKey = 'seckill:stock:' .. voucherId
	local orderKey = 'seckill:order' .. voucherId
	-- 判断是否有库存
	if (tonumber(redis.call('get',stockKey)) <= 0) then
		return 1
	end
	-- 判断用户是否下过单
	if (redis.call('sismember',orderKey,userId) == 1) then
		return 2
	end
	-- 扣库存，吧stockKey的值-1
	redis.call('incrby',stockKey,-1)
	-- 下单，把用户id存到orderKey这个set集合中
	redis.call('sadd',orderKey,userId)
	return 0
	`
)
