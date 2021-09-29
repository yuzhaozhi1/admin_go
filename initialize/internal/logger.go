package internal

import (
	"context"
	"fmt"
	"github.com/yuzhaozhi1/admin_go/global"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}

var (
	// Discard 抛弃
	Discard = New(log.New(ioutil.Discard, "", log.LstdFlags), config{})

	// Default 默认
	Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), config{
		SlowThreshold: 200 * time.Millisecond, // 200 毫秒, 慢sql 的标准
		LogLevel:      logger.Warn,            // sql的报价级别
		Colorful:      true,
	})
	// Recorder 跟踪记录
	Recorder = traceRecorder{Interface:Default, BeginAt: time.Now()}
)

// New  https://github.com/go-gorm/gorm/blob/master/logger/logger.go
// 参考的这个
func New(writer logger.Writer, config config) logger.Interface {
	// 日志的等级,格式化
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}

	return &_logger{
		Writer:       writer,
		config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}

}

// 为了实现 logger.Interface 接口
type _logger struct {
	config
	logger.Writer
	infoStr, warnStr, errStr            string
	traceStr, traceWarnStr, traceErrStr string
}

// LogMode log mode
func (log *_logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *log
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (log *_logger) Info(ctx context.Context, message string, data ...interface{}) {
	if log.LogLevel >= logger.Info {
		log.Printf(log.infoStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn message
func (log *_logger) Warn(ctx context.Context, message string, data ...interface{}) {
	if log.LogLevel >= logger.Warn {
		log.Printf(log.warnStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error message
func (log *_logger) Error(ctx context.Context, message string, data ...interface{}) {
	if log.LogLevel >= logger.Warn {
		log.Printf(log.errStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (log *_logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	/*
		const (
			Silent LogLevel = iota + 1
			Error
			Warn
			Info
		)
	*/
	if log.LogLevel > 0 {
		elapsed := time.Since(begin)
		sql, rows := fc()

		switch {
		// error 级别的 trace
		case err != nil && log.LogLevel >= logger.Error:
			if rows == -1 {
				log.Printf(log.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				log.Printf(log.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		// Warn 级别的trace
		case elapsed > log.SlowThreshold && log.SlowThreshold != 0 && log.LogLevel >= logger.Warn:
			slowLog := fmt.Sprintf("SLOW SQL >= %v", log.SlowThreshold)
			if rows == -1 {
				log.Printf(log.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				log.Printf(log.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		// info 级别的 trace
		case log.LogLevel >= logger.Info:
			if rows == -1 {
				log.Printf(log.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				log.Printf(log.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

// Printf 实现 _logger 里面的logger.Writer接口 中的 Printf 函数
func (log *_logger) Printf(message string, data ...interface{}) {
	if global.GLOBAL_CONFIG.Mysql.LogZap {
		global.GLOBAL_LOG.Info(fmt.Sprintf(message, data...))
	} else {
		log.Writer.Printf(message, data...)
	}
}

// traceRecorder 跟踪记录
type traceRecorder struct {
	logger.Interface
	BeginAt      time.Time // 开始时间
	SQL          string
	RowsAffected int64 // 受影响的行
	Err          error
}

// New 实例化 跟踪记录
func (t traceRecorder) New() *traceRecorder {
	return &traceRecorder{
		Interface: t.Interface,
		BeginAt:   time.Now(),
	}
}

func (t *traceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	t.BeginAt = begin
	t.SQL, t.RowsAffected = fc()
	t.Err = err
}
