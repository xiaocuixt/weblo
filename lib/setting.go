//https://zhuanlan.zhihu.com/p/292253885
package lib

import (
  "fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

func Run() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
  // fmt.Println(viper.GetString("name"))
	// viper解析配置文件
	err := viper.ReadInConfig() 
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

