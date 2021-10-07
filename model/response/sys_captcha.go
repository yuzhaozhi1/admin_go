package response

// SysCaptchaResponse 验证码返回数据结构体
type SysCaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath string `json:"picPath"`
}
