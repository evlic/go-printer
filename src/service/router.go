package service

import (
	"evlic/go-printer/src/service/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	// 开启跨域支持
	Cors(r)
	api := r.Group("/api")
	{
		api.GET("/files", controller.AvailableFiles)
		api.PUT("/job", controller.Upload)
		api.GET("/job", controller.JobInfo)
		api.DELETE("/job/:id", controller.Upload)
		api.GET("/job/:id", controller.Upload)
	}
	InitAdmin(r)
}

// Cors 开启跨域支持
func Cors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", "range")
	r.Use(cors.New(config))
}

func InitAdmin(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("/printer")
		admin.POST("/set")
	}
}
