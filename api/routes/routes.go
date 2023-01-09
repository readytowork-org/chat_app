package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewFirebaseRoutes),
	fx.Provide(NewCrudRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	firebaseRoutes FirebaseRoutes,
	crudRoutes CrudRoutes,
) Routes {
	return Routes{
		firebaseRoutes,
		crudRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
