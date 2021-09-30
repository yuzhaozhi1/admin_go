package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/yuzhaozhi1/admin_go/config"
	"github.com/yuzhaozhi1/admin_go/utils/timer"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	// GLOBAL_DB 全局的 mysql 连接
	GLOBAL_DB    *gorm.DB

	// GLOBAL_REDIS 全局的redis 连接
	GLOBAL_REDIS *redis.Client

	// GLOBAL_CONFIG 全局配置映射
	GLOBAL_CONFIG config.Server

	// GLOBAL_VIPER 配置加载
	GLOBAL_VIPER *viper.Viper

	// GLOBAL_LOG *oplogging.logger
	// GLOBAL_LOG 全局的日志
	GLOBAL_LOG  *zap.Logger

	// GLOBAL_TIMER 全局的定时器任务对象
	GLOBAL_TIMER  timer.Timer = timer.NewTimerTask()

	// GLOBA_Concurrency_Control singleflight 包主要是用来做并发控制, 实现资源访问合并
	GLOBA_Concurrency_Control = &singleflight.Group{}
)
