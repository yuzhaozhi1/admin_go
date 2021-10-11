package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3" // gorm的适配器
	"github.com/yuzhaozhi1/admin_go/global"
	"go.uber.org/zap"
	"strings"
	"sync"
)

// Casbin

var (
	syncedEnforcer *casbin.SyncedEnforcer // 同步执行器
	once           sync.Once
)

// Casbin 初始化casbin 服务, 返回casbin 执行器
func Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.GLOBAL_DB) // 将数据保存到 数据库, 通过适配器
		syncedEnforcer, _ := casbin.NewSyncedEnforcer(global.GLOBAL_CONFIG.Casbin.ModelPath, a)
		syncedEnforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
	})
	err := syncedEnforcer.LoadPolicy()
	if err != nil {
		global.GLOBAL_LOG.Error("加载casbin 服务失败", zap.Any("err", err))
	}
	return syncedEnforcer
}

// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1, ok := args[0].(string)
	if !ok {
		return false, nil
	}

	name2, ok := args[1].(string)
	if !ok {
		return false, nil
	}

	return ParamsMatch(name1, name2), nil

}

// ParamsMatch 自定义规则函数
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用 casbin 的keyMatch2
	// KeyMatch2 判断key1是否与key2的模式匹配(类似于REST式路径)
	// “/foo/bar”匹配“/foo/*” / "/resource1" matches "/:resource"
	return util.KeyMatch2(key1, key2)
}
