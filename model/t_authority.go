package model

import "time"

// SysAuthority 权限表
type SysAuthority struct {
	CreateAt      time.Time      // 创建时间
	UpdateAt      time.Time      // 更新时间
	DeleteAt      *time.Time     `gorm:"index"`
	AuthorityId   string         `json:"authorityId" gorm:"not null;unique;primaryKey;comment:角色ID;size:90"` // 角色id
	AuthorityName string         `json:"authorityName" gorm:"comment:角色名"`                                   // 角色名
	ParentId      string         `json:"parentId" gorm:"comment:父角色ID"`
	DataAuthority []SysAuthority `json:"dataAuthority" gorm:"many2many:sys_data_authority_id"` // 会创建出第三张表  sys_data_authority_id
	Children      []SysAuthority `json:"children" gorm:"-"`
	DefaultRouter string         `json:"defaultRouter" gorm:"default:dashboard; comment:默认菜单"` // 默认菜单(默认 dashboard)
}
