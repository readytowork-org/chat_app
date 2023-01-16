package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewThreadRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	thread ThreadRoutes,
) Routes {
	return Routes{
		thread,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
