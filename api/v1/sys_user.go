package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/middleware"
	"github.com/yuzhaozhi1/admin_go/model"
	"github.com/yuzhaozhi1/admin_go/model/request"
	"github.com/yuzhaozhi1/admin_go/model/response"
	"github.com/yuzhaozhi1/admin_go/service"
	"github.com/yuzhaozhi1/admin_go/utils"
	"go.uber.org/zap"
	"time"
)

// Register 用户注册账号
// data 参数 "用户名 昵称, 密码, 角色ID"
func Register(c *gin.Context) {
	var r request.Register
	_ = c.ShouldBindJSON(&r)
	err := utils.Verify(r, utils.Register)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := model.SysUser{
		Username:    r.Username,
		NickName:    r.NickName,
		Password:    r.Password,
		HeaderImg:   r.HeaderImg,
		AuthorityId: r.AuthorityId,
	}
	// 用户注册
	err, userObj := service.Register(user)
	if err != nil {
		global.GLOBAL_LOG.Error("用户注册失败!", zap.Any("err:", err))
		response.FailWithDetailed(response.SysUserResponse{User: userObj}, "注册失败:"+err.Error(), c)
		return
	} else {
		response.FailWithDetailed(response.SysUserResponse{User: userObj}, "注册成功", c)
	}
}

// Login 用户登录
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var l request.Login
	// 将数据映射到 结构体
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验用户传入的验证码是否正确
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		// 校验用户名和密码
		u := model.SysUser{Username: l.Username, Password: l.Password}
		if err, user := service.Login(&u); err != nil {
			global.GLOBAL_LOG.Error("用户登录失败! 用户名不存在或者密码错误!", zap.Any("err", err))
			response.FailWithMessage("用户登录失败! 用户名不存在或者密码错误!", c)
			return
		} else {
			// 生成jwt token
			fmt.Println(user)
		}
	} else {
		response.FailWithMessage("验证码不正确!", c)
	}
}

// tokenNext 生成token, 用户登录成功后签发token
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.GLOBAL_CONFIG.Jwt.SigningKey)}  // 唯一签名
	claims := request.CustomClaims{
		UUID: user.UUID,
		ID: user.ID,
		NickName: user.NickName,
		Username: user.Username,
		AuthorityId: user.AuthorityId,
		// 缓存时间一天, 缓冲时间内会获得新的token 刷新令牌, 此时一个用户会存在两个有效令牌, 但是前端 只会保留一个,另一个会丢弃
		BufferTime: global.GLOBAL_CONFIG.Jwt.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,  // 签名生效的时间
			ExpiresAt: time.Now().Unix() + global.GLOBAL_CONFIG.Jwt.ExpiresTime,  // 过期时间
			Issuer: "yuZz",
		},
	}
	// 生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GLOBAL_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 是否开启了单点登录
	if !global.GLOBAL_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User: user,
			Token: token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := service.


}

// ChangePassword 用户修改密码
func ChangePassword(c *gin.Context) {

}
