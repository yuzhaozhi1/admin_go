package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	// GLOBAL_DB 全局的 mysql 连接
	GLOBAL_DB    *gorm.DB
	// GLOBAL_REDIS 全局的redis 连接
	GLOBAL_REDIS *redis.Client
	// GLOBAL_VIPER 配置加载
	GLOBAL_VIPER *viper.Viper
	// GLOBA_Concurrency_Control singleflight 包主要是用来做并发控制, 实现资源访问合并
	GLOBA_Concurrency_Control = &singleflight.Group{}
)
