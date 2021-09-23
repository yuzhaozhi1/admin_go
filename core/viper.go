package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yuzhaozhi1/admin_go/global"
	"github.com/yuzhaozhi1/admin_go/utils"
	"os"
	"path/filepath"
)

// Viper 配置文件解析
func Viper(path ...string) *viper.Viper {
	var config string
	// 判断在加载配置,初始化的时候有没有传入 配置文件的url, 如果有传入配置文件的path,就使用传入的path
	if len(path) == 0 {
		// 从运行命令中获取配置文件的path
		flag.StringVar(&config, "c", "","请输入配置文件的path")
		flag.Parse()
		if config == "" {
			// Getenv检索由键命名的环境变量的值。它返回该值，如果该变量不存在，
			// 该值将为空。要区分空值和未设置的值，请使用LookupEnv。
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == ""{
				config = utils.ConfigFile
				fmt.Println("您正在使用默认的配置文件启动, path为:", config)
			}else {
				config = configEnv
				fmt.Println("您正在使用 env 环境变量中的配置文件, path为:", config)
			}
		}else {
			fmt.Println("您正在使用 命令行参数-c 设置的值, path为:", config)
		}
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

	v.OnConfigChange(func(e fsnotify.Event) {  // 如果配置文件发生了变更
		fmt.Println("config file changed:", e.Name)

		// 将配置数据解析到 GLOBAL_CONFIG 中
		if err = v.Unmarshal(&global.GLOBAL_CONFIG); err !=nil {
			fmt.Println("配置文件解析失败,err:", err)
		}
	})

	// 将配置数据解析到 GLOBAL_CONFIG 中
	if err = v.Unmarshal(&global.GLOBAL_CONFIG); err !=nil {
		fmt.Println("配置文件解析失败,err:", err)
	}
	global.GLOBAL_CONFIG.Autocode.Root, _ = filepath.Abs("..")
	// fmt.Println(global.GLOBAL_CONFIG.Mysql.Dsn())
	return v
}
