package model

import (
	"github.com/satori/go.uuid"
	"github.com/yuzhaozhi1/admin_go/global"
)

// SysUser 用户表
type SysUser struct {
	global.BaseModel
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`    // 用户UUID
	Username string    `json:"userName" gorm:"comment:用户登录名;index"` // 用户登录名
	Password string    `json:"-" gorm:"comment:用户密码"`
	NickName string    `json:"nickName" gorm:"default:默认用户昵称;comment:用户昵称"`
	// 用户头像
	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	SideMode    string       `json:"sideMode" gorm:"default:dark;comment:用户侧边栏的主题"`
	ActiveColor string       `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`
	BaseColor   string       `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`
}
