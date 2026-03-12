package conf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config结构体与yaml的配置结构对应
type Config struct {
	Server struct {
		Port int
		Name string
	}

	Database struct {
		DSN string
	}
}

// 实例化一个配置结构体
var AppConfig Config

func InitConfig() {
	//获取当前的环境，默认为dev
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	//读取基本配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs") //.表示运行时的工作目录

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取基础配置文件失败： %w", err)) //%w 包装错误
	}

	//读取特定环境配置并合并同名配置
	viper.SetConfigName("config." + env)
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("未找到特定环境配置文件 config.%s.yaml,将使用默认配置\n", env)
	}

	//将配置映射到配置结构体
	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("解析配置失败: %w", err))
	}

	fmt.Printf("当前运行环境: [%s], 服务器端口: [%d]\n", env, AppConfig.Server.Port)
}
