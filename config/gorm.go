package config

import "fmt"

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             // 服务器地址:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // 高级配置
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // 用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 最大空闲连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启gorm 全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // 是否通过zap 写入日志文件
}

func (m *Mysql) Dsn() string {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", m.Username, m.Password, m.Path, m.Dbname, m.Config)
}
