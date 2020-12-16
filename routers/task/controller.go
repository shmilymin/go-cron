package task

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-cron/models"
	u "go-cron/pkg/util"
	"log"
)

const task_ids_key = "task_ids"

func Router(r *gin.RouterGroup) {
	// 新增
	r.POST("", func(c *gin.Context) {
		task := &models.Task{}
		if e := c.Bind(task); e != nil {
			log.Println(e)
		}
		id := u.UUID()
		task.Id = id
		log.Println(task)
		json, e := json.Marshal(task)
		if e != nil {
			log.Println(e)
		}
		u.Set([]byte(id), json)
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
	// 修改
	r.PUT("", func(c *gin.Context) {

	})
	// 查询全部
	r.GET("", func(c *gin.Context) {

	})
	// 删除
	r.DELETE(":id", func(c *gin.Context) {

	})
}
