package main

import (
	"BiliBili.com/cache"
	"BiliBili.com/model"
	"BiliBili.com/routes"
	"BiliBili.com/service"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

func main() {
	InitConfig()
	println("version:" + viper.GetString("server.version"))
	//初始化Redis
	cache.Redis()
	//初始化数据库
	model.InitDB()
	//创建定时任务
	service.CronJob()
	gin.SetMode(DebugMode)
	r := routes.NewRouter()
	port := viper.GetString("server.port")
	_ = r.Run(":" + port)


}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
