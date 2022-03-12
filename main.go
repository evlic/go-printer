package main

import (
	"evlic/go-printer/src/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
)

func init() {
	// bootstrap.InitLog()
	bootstrap.InitConf()
	bootstrap.InitCron()
	bootstrap.InitCache()
}

func main() {
	// template目录下拥有前端用的一系列静态资源
	box := packr.NewBox("./templates")

	// 映射静态资源文件
	r := gin.Default()
	r.StaticFS("/web", box)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
