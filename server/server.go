package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nanhao/highgin/handler"
	"github.com/nanhao/highgin/env"
	"log"
)

type RouteSetterType func (engine *gin.Engine, w HdlWrapperType)

type AppServer struct {
	engine *gin.Engine
	routeSetter RouteSetterType
	errHandler handler.ErrHandlerType
}

type HdlWrapperType func (handler.Hdl) gin.HandlerFunc

func New(routeSetter RouteSetterType) *AppServer {
	env.Load()

	engine := gin.Default()

	server := &AppServer {
		engine: engine,
		routeSetter: routeSetter,
		errHandler: handler.ErrHandler,
	}

	routeSetter(engine, server.W)

	return server
}

func (ins *AppServer) SetErrHandler(errHandler handler.ErrHandlerType) {
	ins.errHandler = errHandler
}

func (ins *AppServer) Run() {
	listen, err := env.GetStrict("LISTEN")
	if err != nil {
		log.Fatal(err)
	}
	ins.engine.Run(listen)
}

func (ins *AppServer) W(hdl handler.Hdl) gin.HandlerFunc {
	return func (c *gin.Context) {
		ctx := handler.NewCtx(c, ins.errHandler)
		hdl(ctx)
	}
}