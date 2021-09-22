package main

import (
	"BiliBili.com/conf"
	"BiliBili.com/routes"
	"BiliBili.com/service"
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
	conf.InitConfig()
	println("version:" + viper.GetString("server.version"))
	//初始化Redis
	service.CronJob()
	gin.SetMode(DebugMode)
	r := routes.NewRouter()
	port := viper.GetString("server.port")
	_ = r.Run(":" + port)
}
