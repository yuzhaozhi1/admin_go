package service

import (
	"errors"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/model"
	"gorm.io/gorm"
	"time"
)

// JWT 黑名单相关的操作

// JoinInBlackList 将用户的jwtToken 加到redis 中
func JoinInBlackList(jwtList model.JwtBlackList) error {
	err := global.GLOBAL_DB.Create(&jwtList).Error
	return err
}

// IsInBlackList 校验JWT 是否在黑名单中
func IsInBlackList(token string) bool {
	var blackJWT model.JwtBlackList
	return !errors.Is(global.GLOBAL_DB.Where("jwt = ?", token).First(&blackJWT).Error, gorm.ErrRecordNotFound)
}

// GetJWTTokenByRedis 从redis 中获取用户的token
func GetJWTTokenByRedis(userName string) (redisJwt string, err error) {
	redisJwt, err = global.GLOBAL_REDIS.Get(userName).Result()
	return redisJwt, err
}

// SaveJWTTokenToRedis 保存用户的jwt token 到redis中
func SaveJWTTokenToRedis(jwtToken string, userName string) (err error) {
	timer := time.Duration(global.GLOBAL_CONFIG.Jwt.ExpiresTime) * time.Second
	err = global.GLOBAL_REDIS.Set(userName, jwtToken, timer).Err()
	return err
}
