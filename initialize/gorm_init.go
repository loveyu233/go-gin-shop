package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// GormInit The detailed information:
// @Title GormInit
// @Description gorm初始化
// @Param dsn 连接信息
// @Return *gorm.DB
func GormInit(dsn string) *gorm.DB {
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		panic("gorm初始化失败：" + err.Error())
		return nil
	}
	return d
}
