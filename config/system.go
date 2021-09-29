package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 更改为“develop”以跳过开发模式的身份验证
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"`                              // 后端端口
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // 使用的数据库 默认mysql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // 可以指定上传头像的oss
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 单点登录,默认为关闭
}
