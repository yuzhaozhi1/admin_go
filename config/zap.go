package config

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 日志的等级
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 删除的格式
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志的前缀
	Director      string `mapstructure:"director" json:"director" yaml:"director"`                  // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"show-line"`                // 日志显示行号
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}
