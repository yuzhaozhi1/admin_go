package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/yuzhaozhi1/admin_go/global"
	"go.uber.org/zap"
)

// redis redis 配置

// Redis Redis 初始化
func Redis(){
	redisConfig := global.GLOBAL_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr: redisConfig.Addr,
		Password: redisConfig.Password,
		DB: redisConfig.DB,
	})
	// 测试连接redis 是否成功
	pong,err := client.Ping().Result()
	if err != nil {
		global.GLOBAL_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
		return
	}
	global.GLOBAL_LOG.Info("redis connect ping response:", zap.Any("msg:", pong))

	global.GLOBAL_REDIS = client
	fmt.Println("连接到redis 成功")
}
