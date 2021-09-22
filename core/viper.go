package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper 配置文件解析
func Viper(path ...string) *viper.Viper {
	var config string
	// 判断在加载配置,初始化的时候有没有传入 配置文件的url, 如果有传入配置文件的path,就使用传入的path
	if len(path) == 0 {

	} else {
		config = path[0]
		fmt.Printf("你使用的为 初始化core.Viper() 函数时传入的 配置文件, 路径为%s\n", config)
	}
	// 先获取一个新的 Viper 指针
	v := viper.New()
	// 设置 配置文件的 路径
	v.SetConfigFile(config)
	v.SetConfigType("yaml") // 设置配置文件的类型为 yaml
	err := v.ReadInConfig() // 将配置文件进行读取, 并以key:value 的格式进行 储存
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败,err:%s\n", err))
	}

	// 监视配置文件的更新, 哨兵模式
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.)

	})

}
