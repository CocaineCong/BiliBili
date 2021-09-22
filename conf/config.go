package conf

import (
	"BiliBili.com/cache"
	"BiliBili.com/model"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	fmt.Println("workDir",workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//初始化数据库
	cache.Redis()
	model.InitDB()
	//创建定时任务
}
