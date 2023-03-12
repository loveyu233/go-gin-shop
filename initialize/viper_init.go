package initialize

import (
	"github.com/spf13/viper"
	"go-gin-shop/global"
)

// Init_viper The detailed information:
// @Title Init_viper
// @Description viper配置信息
func Init_viper() {
	v := viper.New()
	v.SetConfigFile(global.ConfigFilePath)
	err := v.ReadInConfig()
	if err != nil {
		panic("viper初始化错误：" + err.Error())
		return
	}
	err = v.Unmarshal(&global.Config)
	if err != nil {
		panic("mysql redis 初始化错误：" + err.Error())
		return
	}
}
