package request

// 用户相关的请求参数校验

// Register 用户注册结构体
type Register struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://qmplusimg.henrongyi.top/head.png'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"` // auth id
}

type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码id
}
