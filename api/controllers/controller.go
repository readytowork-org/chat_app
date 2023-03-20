package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewRoomController),
	fx.Provide(NewMessageController),
	fx.Provide(NewThreadController),
	fx.Provide(NewJwtAuthController),
)
