package config

type MysqlConfig struct {
	Username string `json:"username,omitempty" yaml:"username"`
	Password string `json:"password,omitempty" yaml:"password"`
	Addr     string `json:"addr,omitempty" yaml:"addr"`
	Config   string `json:"config,omitempty" yaml:"config"`
	DbName   string `json:"dbName,omitempty" yaml:"dbName"`
}

func (m *MysqlConfig) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Addr + ")/" + m.DbName + "?" + m.Config
}
