package service

import (
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/model"
	"github.com/yuzhaozhi1/admin_go/utils"
)

// 用户相关的服务类

// Login 用户登录,校验用户名和密码的函数
func Login(u *model.SysUser)(err error, userInfo *model.SysUser){
	var user model.SysUser
	// 获取加密后的密码
	u.Password = utils.MD5V([]byte(u.Password))
	// global.GLOBAL_DB.Where("username = ? and password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	err = global.GLOBAL_DB.Where("username = ? and password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
