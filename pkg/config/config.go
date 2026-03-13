package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config结构体与yaml的配置结构对应
// Server 表示程序后台
// Database 表示数据库
type Config struct {
	Server   ServerConfig   `mapstructre:"server"`
	Database DatabaseConfig `mapstructre:"database"`
}

type ServerConfig struct {
	Port int
	Name string
}

type DatabaseConfig struct {
	DSN string
}

// 实例化一个配置结构体
var AppConfig Config

// 根据环境变量，初始化配置
func InitConfig() {
	//获取当前的环境，默认为dev
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	//切换到项目的根目录下面
	Init()

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

// 将目录切换到项目的根目录下去
func Init() {
	for range 10 {
		if _, err := os.Stat("go.mod"); err == nil {
			// 找到 go.mod
			return
		}
		//没找到，往上走一层
		if err := os.Chdir(".."); err != nil {
			panic("无法切换目录:" + err.Error())
		}
	}
}
