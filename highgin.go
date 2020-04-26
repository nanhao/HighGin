package highgin

import (
	"github.com/nanhao/highgin/server"
)

var app *server.AppServer
var appCreated bool = false

func NewApp(routeSetter RouteSetterType) *server.AppServer {
	if !appCreated {
		app = server.New(routeSetter)
	}
	return app
}

func App() *server.AppServer {
	return app
}