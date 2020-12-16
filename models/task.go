package models

type Task struct {
	Id         string `form:"id" json:"id"`
	AppName    string `form:"app_name" json:"app_name"`
	TaskName   string `form:"task_name" json:"task_name"`
	Cron       string `form:"cron" json:"cron"`
	CreateTime string `form:"create_time" json:"create_time"`
	UpdateTime string `form:"update_time" json:"update_time"`
}
