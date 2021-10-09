package service

import (
	"errors"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/model"
	"gorm.io/gorm"
)

// JWT 黑名单相关的操作

// IsInBlackList 校验JWT 是否在黑名单中
func IsInBlackList(token string) bool {
	var blackJWT model.JwtBlackList
	return !errors.Is(global.GLOBAL_DB.Where("jwt = ?", token).First(&blackJWT).Error, gorm.ErrRecordNotFound)
}
