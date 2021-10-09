package model

import "github.com/yuzhaozhi1/admin_go/global"

type JwtBlackList struct {
	global.BaseModel
	jwt string `gorm:"type:text;comment:jwt"`
}
