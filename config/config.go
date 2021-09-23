package config

// Server 项目总的配置类
type Server struct {
	// 日志
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// auto
	Autocode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
