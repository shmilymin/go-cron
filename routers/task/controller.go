package task

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-cron/models"
	u "go-cron/pkg/utils"
	"log"
)

func Router(r *gin.RouterGroup) {
	// 新增or修改
	r.POST("", func(c *gin.Context) {
		task := &models.Task{}
		if e := c.Bind(task); e != nil {
			log.Println(e)
		}
		now := u.Now()
		task.CreateTime = now
		task.UpdateTime = now
		if task.Id == "" {
			task.Id = u.UUID()
		}
		log.Println(task)

		u.Ok(c)
	})
	// 根据id查找
	r.GET(":id", func(c *gin.Context) {
		id := c.Param("id")
		val, err := u.Get([]byte(id))
		if err != nil {
			log.Println(err)
		}
		task := models.Task{}
		if err := json.Unmarshal(val, &task); err != nil {
			log.Println(err)
		}
		u.OkData(c, task)
	})
	// 查询全部
	r.GET("", func(c *gin.Context) {
		p := &u.Page{}
		c.Bind(&p)

	})
	// 删除
	r.DELETE(":id", func(c *gin.Context) {
		id := c.Param("id")
		if u.Del([]byte(id)) {
			u.Ok(c)
		}
		u.Err(c)
	})
}
