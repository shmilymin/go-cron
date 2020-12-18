package task

import (
	"github.com/gin-gonic/gin"
	"go-cron/models"
	u "go-cron/pkg/utils"
	"log"
	"strconv"
)

func Router(r *gin.RouterGroup) {
	// 新增or修改
	r.POST("", func(c *gin.Context) {
		task := &models.Task{}
		if e := c.Bind(task); e != nil {
			log.Println(e)
		}
		now := u.NowTime()
		task.CreateTime = now
		task.UpdateTime = now
		log.Println(task)
		if task.Id == 0 {
			task.Save()
		} else {
			task.Update()
		}
		u.Ok(c)
	})
	// 查询全部
	r.GET("", func(c *gin.Context) {
		p := &u.Page{}
		c.Bind(&p)
		task := &models.Task{}
		c.Bind(&task)
		tasks := models.Tasks{}
		tasks.List(p)
		count := task.Count()
		u.OkData(c, map[string]interface{}{"tasks": tasks, "count": count})
	})
	// 删除
	r.DELETE(":id", func(c *gin.Context) {
		task := &models.Task{}
		task.Id, _ = strconv.ParseUint(c.Param("id"), 10, 64)
		task.Del()
		u.Ok(c)
	})
}
