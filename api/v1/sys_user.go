package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yuzhaozhi1/admin_go/model/request"
	"github.com/yuzhaozhi1/admin_go/model/response"
	"github.com/yuzhaozhi1/admin_go/utils"
)



// Login 用户登录
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context){
	var l request.Login
	// 将数据映射到 结构体
	err := c.ShouldBindJSON(&l)
	if err != nil{
		response.FailWithMessage(err.Error(), c)
		return
	}

	utils.Verify(l, utils.LoginVerify)




}
