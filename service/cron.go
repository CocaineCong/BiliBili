package service

import (
	"BiliBili.com/pkg/utils"
	"github.com/robfig/cron"
)

var Cron *cron.Cron

func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}

	//每天凌晨1点执行
	err := Cron.AddFunc("0 0 1 * * ?", ClicksStoreInDB)
	if err != nil {
		utils.Logfile("[Error]", " cornJob error  "+err.Error())
	}
	Cron.Start()
	println("created cron job")
}
