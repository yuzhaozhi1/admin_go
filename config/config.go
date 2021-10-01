package config

// Server 项目总的配置类
type Server struct {

	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	// 系统配置
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// 日志
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// auto
	Autocode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// Timer 定时器
	Timer

	// 文件储存的几个类型 本地, 七牛
	Local Local `mapstructure:"local" json:"local" yaml:"local"`

}
