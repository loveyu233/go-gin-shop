package config

type System struct {
	Ip   string `json:"ip,omitempty" yaml:"ip"`
	Port string `json:"port,omitempty" yaml:"port"`
}

func (s System) Dsn() string {
	return s.Ip + ":" + s.Port
}
