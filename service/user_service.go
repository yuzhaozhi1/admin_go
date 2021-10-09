package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/model"
	"github.com/yuzhaozhi1/admin_go/utils"
	"gorm.io/gorm"
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

// Register 用户注册的函数
func Register(registerUser model.SysUser)(err error, userObj model.SysUser) {
	var user model.SysUser
	// 校验用名是否被使用
	// errors.Is 报告err链中的任何错误是否与目标匹配, 如果不匹配就返回 false
	if !errors.Is(global.GLOBAL_DB.Where("username = ?", registerUser.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("该用户名已被注册"), userObj
	}
	// 加密用户的密码, 使用MD5 加密
	registerUser.Password = utils.MD5V([]byte(registerUser.Password))
	registerUser.UUID = uuid.NewV4()

	// 保存到数据库
	err = global.GLOBAL_DB.Create(&registerUser).Error
	return err, registerUser
}
