package initialize

import (
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/initialize/internal"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// mysql 数据库链接

func Gorm() *gorm.DB {
	switch global.GLOBAL_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	m := global.GLOBAL_CONFIG.Mysql
	dsn := m.Dsn()
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // Dsn data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度, Mysql 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式, Mysql 5.7 之前的 数据库和 mariaDB 不支持
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列, Mysql 8 之前的数据库和mariaDB 不支持
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		global.GLOBAL_LOG.Error("Mysql 连接异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()

		// https://gorm.io/zh_CN/docs/connecting_to_the_database.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// gormConfig 根据配置决定是否开启日志
func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true} // 迁移时禁用外键约束

	switch global.GLOBAL_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		config.Logger = internal.Default.LogMode(logger.Info)
	}
	return config
}
