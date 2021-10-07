package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/model/response"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// // @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的 driver
	driver := base64Captcha.NewDriverDigit(
		global.GLOBAL_CONFIG.Captcha.ImgHeight, global.GLOBAL_CONFIG.Captcha.ImgWidth, global.GLOBAL_CONFIG.Captcha.KeyLong,
		0, 80)

	cp := base64Captcha.NewCaptcha(driver, store)

	if id, b64s, err := cp.Generate(); err != nil {
		global.GLOBAL_LOG.Error("获取验证码失败!", zap.Any("err:", err))
		response.FailWithMessage("获取验证码失败!", c)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "获取验证码成功!", c)
	}
}
