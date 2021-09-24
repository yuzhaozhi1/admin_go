package utils

import (
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/yuzhaozhi1/admin_go/global"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

// rotate logs  按日期轮询file-rotatelogs

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.GLOBAL_CONFIG.Zap.Director, "%Y-%m-%d.log"),  // 日志文件的路径
		zaprotatelogs.WithLinkName(global.GLOBAL_CONFIG.Zap.LinkName), // 生成软链，指向最新日志文件
		zaprotatelogs.WithMaxAge(7*24*time.Hour),                      // //clear 最小分钟为单位
		// rotate 最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
		zaprotatelogs.WithRotationTime(24*time.Hour), // number 默认7份 大于7份 或到了清理时间 开始清理
	)

	// 判断需不需要在终端输出
	if global.GLOBAL_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
