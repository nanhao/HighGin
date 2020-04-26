package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nanhao/highgin/err"
)

func ErrHandler(err err.IErr, ctx *gin.Context) {
	err.Log()
	ctx.JSON(err.HttpStatus(), gin.H {
		"code": err.ErrCode(),
		"message": err.Error(),
	})
}