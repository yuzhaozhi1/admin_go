package core

import (
	"fmt"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
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
	// 从配置文件中获取日志的level 初始化配置文件的level
	switch global.GLOBAL_CONFIG.Zap.Level {
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
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	// 如果要显示行号
	if global.GLOBAL_CONFIG.Zap.ShowLine {
		// AddCaller将Logger配置为使用ZAP调用者的文件名和行号注释每条消息
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message", // 输入信息的key名
		LevelKey:       "level",   // 输出日志级别的key名
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GLOBAL_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     CustomTimeEncoder,              // 自定义日志输出时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
	}

	// 获取日志的编码
	switch {
	case global.GLOBAL_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.GLOBAL_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.GLOBAL_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.GLOBAL_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取 zapcore.Encoder 编码器
func getEncoder() zapcore.Encoder {
	if global.GLOBAL_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewJSONEncoder(getEncoderConfig())
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GLOBAL_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

// getEncoderCore 返回 zapcore.Core, 用于实例化zap 日志对象
func getEncoderCore() (core zapcore.Core) {
	writer, err := utils.GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get writer syncer err%v\n", err.Error())
		return nil
	}
	// NewCore创建一个Core，将日志写入WriteSyncer。
	return zapcore.NewCore(getEncoder(), writer, level)
}
