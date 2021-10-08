package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)  // 用于优雅的重启http 服务
	s.ReadHeaderTimeout = 10 * time.Millisecond   // 允许读取的最大时间
	// s.WriteTimeout = 10 * time.Millisecond  // 允许写入的最大时间
	s.WriteTimeout = 180 * time.Second
	s.MaxHeaderBytes = 1 << 20   // 请求头的最大字节数
	return s
}

/*
	目的
	不关闭现有连接（正在运行中的程序）
	新的进程启动并替代旧进程
	新的进程接管新的连接
	连接要随时响应用户的请求，当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况

	流程
	替换可执行文件或修改配置文件
	发送信号量 SIGHUP
	拒绝新连接请求旧进程，但要保证已有连接正常
	启动新的子进程
	新的子进程开始 Accet
	系统将新的请求转交新的子进程
	旧进程处理完所有旧连接后正常结束
*/
