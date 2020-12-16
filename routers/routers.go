package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-cron/routers/task"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	setUpConfig(r)
	setUpRouter(r)
	return r
}

// 初始化应用设置
func setUpConfig(r *gin.Engine) {
	// 跨域配置
	r.Use(cors.Default())
}

// 设置路由
func setUpRouter(r *gin.Engine) {
	api := r.Group("/api")
	task.Router(api.Group("/task"))
}
