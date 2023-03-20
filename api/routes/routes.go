package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewThreadRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoomRoutes),
	fx.Provide(NewMessageRoutes),
	fx.Provide(NewObtainJwtTokenRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	thread ThreadRoutes,
	user UserRoutes,
	room RoomRoutes,
	messages MessageRoutes,
	jwt ObtainJwtTokenRoutes,

) Routes {
	return Routes{
		thread,
		user,
		room,
		messages,
		jwt,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
