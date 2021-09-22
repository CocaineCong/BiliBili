package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)


var DB *gorm.DB

func InitDB() {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, path)
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	db.SingularTable(true)       //默认不加复数s
	db.DB().SetMaxIdleConns(20)  	//设置连接池，空闲
	db.DB().SetMaxOpenConns(100) 	//设置打开最大连接
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	//数据库迁移
	migration()
}
