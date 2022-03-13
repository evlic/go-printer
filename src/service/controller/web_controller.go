package controller

import (
	"evlic/go-printer/src/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// web：提交打印任务、查看用户自己提交的打印任务状态。
// 提交任务为：上传、提交到打印机系统

// Upload 上传文件，仅限格式
func Upload(ctx *gin.Context) {
	uploadFolder := conf.Conf.TempDir
	if form, err := ctx.MultipartForm(); err != nil {
		log.Error("上传发生错误", err)
		ctx.String(http.StatusBadRequest, "错误的上传请求")
	} else {
		files := form.File["file"]
		resStatus := gin.H{}
		cnt := 0
		for _, file := range files {
			if err := ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", uploadFolder, file.Filename)); err != nil {
				resStatus[file.Filename] = "error"
				cnt++
			} else {
				resStatus[file.Filename] = "success"
			}
		}
		if cnt == 0 {
			ctx.JSON(http.StatusPartialContent, resStatus)
		} else {
			ctx.String(http.StatusOK, "上传成功")
		}
	}
}

// AvailableFiles 可用文件列表
func AvailableFiles(ctx *gin.Context) {
	
}

// SubmitJob 提交任务
func SubmitJob(ctx *gin.Context) {
	// 前端提交 JSON 表示选中文件的情况，提交打印请求
	
}

// JobCancel 取消任务的执行
func JobCancel(ctx *gin.Context) {

}

// JobInfo 打印任务执行状态
func JobInfo(ctx *gin.Context) {

}
