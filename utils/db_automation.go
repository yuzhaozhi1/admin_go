package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// db_automation  数据库自动化操作

// ClearTable 清理数据库表数据
// db 数据库表对象  *gorm.DB  tableName(表名) string compareFiled(比较字段) string  interval(间隔) string
func ClearTable(db *gorm.DB, tableName string, compareFiled string, interval string) error {
	if db == nil {
		return errors.New("db cannot be empty")
	}
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0, 解析时间小于 0")
	}
	return db.Debug().Exec(fmt.Sprintf(
		"delete from %s where %s < ?", tableName, compareFiled), time.Now().Add(-duration)).Error

}
