package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	// fx.Provide(NewMiddewares),
	fx.Provide(NewFirebaseAuth),
	fx.Provide(NewJWTAuthMiddleWare),
)
