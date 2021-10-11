package model

import "github.com/yuzhaozhi1/admin_go/global"

type JwtBlackList struct {
	global.BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
