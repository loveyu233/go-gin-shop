package config

type Config struct {
	Mysql  MysqlConfig `yaml:"mysql"`
	Redis  RedisConfig `yaml:"redis"`
	System System      `yaml:"system"`
	Mongo  MongoConfig `yaml:"mongodb"`
}
