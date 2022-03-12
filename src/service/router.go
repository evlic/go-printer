package service

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 开启跨域支持
func Cors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", "range")
	r.Use(cors.New(config))
}

func InitApi(r *gin.Engine) {
	Cors(r)
	api := r.Group("/api")
	jobs := api.Group("/jobs")
	{
		// 打印机的任务列表
		jobs.GET("/:printer")
		// 打印机的任务详细
		jobs.GET("/:printer/:id")
		// 新增任务
		jobs.PUT("/:printer")
		// 删除任务
		jobs.DELETE("/:printer/:id")
	}
	printer := api.Group("p")
	{
		// 可用打印机列表
		printer.GET("/available")
		printer.GET("/info/:printer")

		// 修改默认打印机
		printer.PUT("/:printer")
		
		// 新增打印机
		printer.POST("/add")
		// 获得可用打印机连接
		printer.GET("/links")
	}
}
