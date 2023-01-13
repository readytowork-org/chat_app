package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewFirebaseRoutes),
	fx.Provide(NewCrudRoutes),
	fx.Provide(NewThreadRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	firebaseRoutes FirebaseRoutes,
	crudRoutes CrudRoutes,
	thread ThreadRoutes,

) Routes {
	return Routes{
		firebaseRoutes,
		crudRoutes,
		thread,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
