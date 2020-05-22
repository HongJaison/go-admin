package example

import (
	c "github.com/HongJaison/go-admin/modules/config"
	"github.com/HongJaison/go-admin/modules/service"
	"github.com/HongJaison/go-admin/plugins"
)

type Example struct {
	*plugins.Base
}

func NewExample() *Example {
	return &Example{
		Base: &plugins.Base{PlugName: "example"},
	}
}

func (e *Example) InitPlugin(srv service.List) {
	e.InitBase(srv)
	e.App = e.initRouter(c.Prefix(), srv)
}
