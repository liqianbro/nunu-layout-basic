package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func NewConfig() *viper.Viper {
	envConf := os.Getenv("APP_CONF")

	// 直接赋值为默认值，避免重复判断
	if envConf == "" {
		flag.StringVar(&envConf, "conf", "config/local.yaml", "config path, eg: -conf config/local.yaml")
		flag.Parse()
	}

	fmt.Println("Loading config file:", envConf)

	// 创建 Viper 实例
	conf := viper.New()
	conf.SetConfigFile(envConf)

	// 读取配置文件
	if err := conf.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	// 实现配置文件实时更新
	conf.WatchConfig()
	conf.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return conf
}
