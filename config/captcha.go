package config

// 验证码配置

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`       // 验证码的长度,有几位数
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`    // 验证码的宽度
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"` // 验证吗的高度
}
