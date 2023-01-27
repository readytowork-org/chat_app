package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewFirebaseService),
	fx.Provide(NewCrudService),
	fx.Provide(NewUserService),
)
