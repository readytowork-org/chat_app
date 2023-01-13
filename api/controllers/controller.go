package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewFirebaseController),
	fx.Provide(NewCrudController),
	fx.Provide(NewThreadController),
)
