package controller

import (
	"evlic/go-printer/src/util"

	"github.com/gin-gonic/gin"
)

func GetAilablePrinter(ctx *gin.Context) {
	util.DoCmd("lpstat", "-t")
}