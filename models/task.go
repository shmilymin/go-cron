package models

import (
	u "go-cron/pkg/utils"
	"log"
)

var tasks []Task

type Task struct {
	Model
	AppName  string `form:"app_name"`
	TaskName string `form:"task_name"`
	Cron     string `form:"cron"`
}

// 新增
func (m *Task) Save() {
	m.Id = uint64(len(tasks) + 1)
	tasks = append(tasks, *m)
}

// 修改
func (m *Task) Update() {
	m = &tasks[m.Id-1 : m.Id][0]
}

// 删除
func (m *Task) Del() {
	tasks = append(tasks[:m.Id-1], tasks[m.Id:]...)
}

// 列表
func (m *Task) List(p *u.Page) (ts []Task) {
	l := len(tasks)
	idx1 := (p.Page - 1) * p.Limit
	idx2 := p.Page * p.Limit
	if idx2 >= l {
		idx2 = l
	}
	list := tasks[idx1:idx2]
	log.Println(list)
	return list
}

// 总数
func (m *Task) Count() int {
	return len(tasks)
}
