package test

import (
	"go-cron/models"
	"go-cron/pkg/utils"
	"testing"
)

func TestSlice(t *testing.T) {
	var tasks []models.Task
	task1 := models.Task{}
	task1.Id = 1
	task1.TaskName = "t1"
	tasks = append(tasks, task1)
	t.Log(tasks)
	task2 := models.Task{}
	task2.Id = 2
	task2.TaskName = "t2"
	tasks = append(tasks, task2)
	t.Log(tasks)
	task := &tasks[1:2][0]
	task.TaskName = "tt2"
	t.Log(tasks)
}

func TestDate(t *testing.T) {
	t.Log(utils.NowDT())
}
