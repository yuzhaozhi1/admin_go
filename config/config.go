package config

// Server 项目总的配置类
type Server struct {

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

}
