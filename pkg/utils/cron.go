package utils

import (
	"github.com/robfig/cron/v3"
	"log"
)

var ServerCron *cron.Cron

// 初始化定时任务
func init() {
	log.Println("=========cron start========")
	ServerCron = cron.New(cron.WithSeconds())
	ServerCron.Start()
	log.Println("=========cron end========")
}
