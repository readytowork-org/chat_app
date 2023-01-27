package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewThreadRoutes),
	fx.Provide(NewUserRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	thread ThreadRoutes,
	user UserRoutes,

) Routes {
	return Routes{
		thread,
		user,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
