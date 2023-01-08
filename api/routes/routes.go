package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewFirebaseRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	firebaseRoutes FirebaseRoutes,
) Routes {
	return Routes{
		firebaseRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
