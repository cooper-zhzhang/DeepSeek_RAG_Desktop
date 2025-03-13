package global

import (
	"github.com/spf13/viper"
)

var DBViper *viper.Viper

func initViper() {
	// 设置配置文件的名称和路径
	viper.SetConfigName("config")   // 配置文件名称（不包含扩展名）
	viper.SetConfigType("yaml")     // 指定配置文件格式为 YAML
	viper.AddConfigPath("./config") // 添加其他可能的路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	DBViper = viper.New()
	DBViper.SetConfigName("db")       // 配置文件名称（不包含扩展名）
	DBViper.SetConfigType("yaml")     // 指定配置文件格式为 YAML
	DBViper.AddConfigPath("./config") // 添加其他可能的路径
	if err := DBViper.ReadInConfig(); err != nil {
		panic(err)
	}

}
