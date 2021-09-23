package core

import (
	"fmt"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	// 判断是否有这个日志文件夹,没有就创建这个文件夹
	if ok, _ := utils.PathExists(global.GLOBAL_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v Director\n", global.GLOBAL_CONFIG.Zap.Director)
		// mkdir创建具有指定名称和权限位的新目录(在umask之前)。
		// 如果有错误，则类型为*PathError
		_ = os.Mkdir(global.GLOBAL_CONFIG.Zap.Director, os.ModePerm)
	}
	switch global.GLOBAL_CONFIG.Zap.Level { // 初始化配置文件的level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn": // 警告
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		looger = zap.New()
	}

}
