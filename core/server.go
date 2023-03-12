package core

import (
	"go-gin-shop/global"
	"go-gin-shop/initialize"
	"log"

	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

// RunServer The detailed information:
// @Title RunServer
// @Description 运行服务
func RunServer() {
	// 读取配置文件
	initialize.Init_viper()

	// gorm初始化
	global.MysqlDb = initialize.GormInit(global.Config.Mysql.Dsn())

	db, _ := global.MysqlDb.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.SetMaxIdleConns(100)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(200)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)

	// redis初始化
	global.RedisDb = initialize.InitRedis(global.Config.Redis.Dsn())

	// mongobd初始化
	global.MongoDb = initialize.InitMongoDb(global.Config.Mongo.Dsn())

	// 初始化casbin
	initialize.InitCasbin()

	// bloomfilter初始化
	global.Bloomfilter = initialize.InitBloomfilter()

	// 路由初始化
	routers := initialize.Routers()
	if routers.Run(global.Config.System.Dsn()) != nil {
		panic("运行失败")
	}
}
