package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func InitServer(address string, c *gin.Engine) server {
	s := endless.NewServer(address, c)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20 // 1024 * 1024
	return s
}

func RunServer()  {

	// 判断是否使用了单点登录, 如果开启了就初始化redis
	if global.GLOBAL_CONFIG.System.UseMultipoint {
		// 连接到redis
		initialize.Redis()
	}


}
