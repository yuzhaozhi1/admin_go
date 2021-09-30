package global

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int            `gorm:"primaryKey;comment:主键" json:"id" form:"id"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"` // 创建时间
	UpdatedAt time.Time      `gorm:"comment:更新时间"`                   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`    // 删除时间
}
