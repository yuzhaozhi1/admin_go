package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/model"
	"github.com/yuzhaozhi1/admin_go/model/request"
	"github.com/yuzhaozhi1/admin_go/model/response"
	"github.com/yuzhaozhi1/admin_go/service"
	"github.com/yuzhaozhi1/admin_go/utils"
	"go.uber.org/zap"
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
	}else {
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
			fmt.Println(user)
		}
	} else {
		response.FailWithMessage("验证码不正确!", c)
	}
}

// ChangePassword 用户修改密码
func ChangePassword(c *gin.Context)  {

}
