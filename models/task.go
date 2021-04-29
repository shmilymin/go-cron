package models

import (
	"github.com/robfig/cron/v3"
	u "go-cron/pkg/utils"
	"log"
)

var tasks []Task

type Task struct {
	Model
	AppName  string `form:"app_name"`
	TaskName string `form:"task_name"`
	Cron     string `form:"cron"`
	Url      string `form:"url"`
	CronId   int    `form:"cron_id"`
}

// AddCorn 新增定时任务
func (m *Task) AddCorn() {
	cronId, _ := u.ServerCron.AddFunc(m.Cron, func() {
		// 请求url
		log.Println("开始执行任务")
	})
	m.CronId = int(cronId)
}

// DelCron 删除定时任务
func (m *Task) DelCron() {
	u.ServerCron.Remove(cron.EntryID(m.CronId))
}

// Save 新增
func (m *Task) Save() {
	m.Id = uint64(len(tasks) + 1)
	m.AddCorn()
	tasks = append(tasks, *m)
}

// Update 修改
func (m *Task) Update() {
	task := &tasks[m.Id-1 : m.Id][0]
	task.DelCron()
	task.TaskName = m.TaskName
	task.Url = m.Url
	task.AppName = m.AppName
	task.Cron = m.Cron
	m.AddCorn()
	task.CronId = m.CronId
}

// Del 删除
func (m *Task) Del() {
	m.DelCron()
	tasks = append(tasks[:m.Id-1], tasks[m.Id:]...)
}

// List 列表
func (m *Task) List(p *u.Page) (ts []Task) {
	l := len(tasks)
	if l == 0 {
		return []Task{}
	}
	idx1 := (p.Page - 1) * p.Limit
	idx2 := p.Page * p.Limit
	if idx2 >= l {
		idx2 = l
	}
	list := tasks[idx1:idx2]
	log.Println(list)
	return list
}

// Count 总数
func (m *Task) Count() int {
	return len(tasks)
}
