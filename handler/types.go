package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nanhao/highgin/err"
)

type Obj map[string]interface{}

type Hdl func (ctx *Ctx)

type ErrHandlerType func (err.IErr, *gin.Context)