package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nanhao/highgin/err"
)

type Ctx struct {
	ctx *gin.Context
	errHandler ErrHandlerType
}

func NewCtx(c *gin.Context, errHandler ErrHandlerType) *Ctx {
	ctx := &Ctx {
		ctx: c,
		errHandler: errHandler,
	}
	return ctx
}

func (ins *Ctx) C() *gin.Context {
	return ins.ctx
}

func (ins *Ctx) Success(obj Obj) {
	ins.ctx.JSON(200, obj)
}

func (ins *Ctx) Fail(err err.IErr) {
	ins.errHandler(err, ins.ctx)
}