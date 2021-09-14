package main
import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/yuzhaozhi1/admin_go/global"
)

func main() {
	// 使用 viper 来加载配置信息
	global.GLOBAL_VIPER = core.Viper()



}
