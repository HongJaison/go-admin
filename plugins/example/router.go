package example

import (
	"github.com/HongJaison/go-admin/context"
	"github.com/HongJaison/go-admin/modules/auth"
	"github.com/HongJaison/go-admin/modules/db"
	"github.com/HongJaison/go-admin/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}