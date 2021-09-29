package main
import (
	"github.com/yuzhaozhi1/admin_go/core"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/initialize"
)

func main() {
	// 使用 viper 来加载配置信息
	global.GLOBAL_VIPER = core.Viper()
	// 初始化zap 日志库
	global.GLOBAL_LOG = core.Zap()
	global.GLOBAL_DB = initialize.Gorm()

}
