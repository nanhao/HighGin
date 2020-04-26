package highgin

import (
	"github.com/nanhao/highgin/server"
	"github.com/nanhao/highgin/handler"
	"github.com/nanhao/highgin/err"
)

type Ctx = handler.Ctx
type Obj = map[string]interface{}
type RouteSetterType = server.RouteSetterType
type ErrorCode = err.ErrorCode
type Err = err.Err
type IErr = err.IErr
type AppServer = server.AppServer
type HdlWrapperType = server.HdlWrapperType