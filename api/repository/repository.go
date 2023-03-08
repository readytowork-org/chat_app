package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewFirebaseRepository),
	fx.Provide(NewUserRepository),
	fx.Provide(NewRoomRepository),
	fx.Provide(NewMessageRepository),
)
