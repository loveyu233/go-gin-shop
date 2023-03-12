package config

import "fmt"

type RedisConfig struct {
	Password string `json:"password" yaml:"password"`
	Addr     string `json:"addr" yaml:"addr"`
	Db       int    `json:"db" yaml:"db"`
}

func (r *RedisConfig) Dsn() string {
	//	"redis://<user>:<pass>@localhost:6379/<db>"
	return fmt.Sprintf("redis://root:%s@%s/%d", r.Password, r.Addr, r.Db)
}
