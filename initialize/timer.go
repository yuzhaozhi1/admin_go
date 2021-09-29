package initialize

import (
	"fmt"
	"github.com/yuzhaozhi1/admin_go/config"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/utils"
)

// timer 定时器

func Timer() {
	if global.GLOBAL_CONFIG.Timer.Start {
		for _, detail := range global.GLOBAL_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GLOBAL_TIMER.AddTaskByFunc("ClearDB", global.GLOBAL_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GLOBAL_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("time error err:", err)
					}
				})
			}(detail)
		}
	}
}
