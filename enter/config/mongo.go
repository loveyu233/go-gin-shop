package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type MongoConfig struct {
	Dbname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Param    string `yaml:"param"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (MongoConfig) Dsn() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/?%s",
		viper.GetString("mongo.username"),
		viper.GetString("mongo.password"),
		viper.GetString("mongo.host"),
		viper.GetString("mongo.port"),
		viper.GetString("mongo.param"),
	)
}
