package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewFBApp),
	fx.Provide(NewEnv),
	fx.Provide(NewLogger),
	fx.Provide(NewRouter),
	fx.Provide(NewDatabase),
)
