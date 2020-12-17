package models

import (
	"encoding/json"
	u "go-cron/pkg/utils"
	"log"
)

type Task struct {
	Id         string `form:"id" json:"id"`
	AppName    string `form:"app_name" json:"app_name"`
	TaskName   string `form:"task_name" json:"task_name"`
	Cron       string `form:"cron" json:"cron"`
	CreateTime string `form:"create_time" json:"create_time"`
	UpdateTime string `form:"update_time" json:"update_time"`
}

const ids = "task-ids"

func (m *Task) addCache() error {
	ids, e := u.Get([]byte(ids))
	if e != nil {
		return e
	}
	var taskIds []string
	if e := json.Unmarshal(ids, &taskIds); e != nil {
		return e
	}
	taskIds = append(taskIds, m.Id)
	ids, e = json.Marshal(taskIds)
	if e != nil {
		return e
	}

	// 保存定时任务
	json, e := json.Marshal(m)
	if e != nil {
		log.Println(e)
	}
	if e = u.Set([]byte(m.Id), json); e != nil {
		log.Println(e)
	}
}
