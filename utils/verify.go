package utils

// verify  验证

var (
	// LoginVerify 登录参数验证
	LoginVerify = Rules{"CaptchaId": {NotEmpty()}, "Captcha":{NotEmpty()}, "Username":{NotEmpty()}, "Password":{NotEmpty()}}
)
