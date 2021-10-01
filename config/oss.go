package config

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件储存的位置
}

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                // 储存的区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                          // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`                     // CDN 加速域名
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`                  // 是否使用https
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`               // 密钥 AccessKey
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`               // 密钥 SecretKey
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速
}
