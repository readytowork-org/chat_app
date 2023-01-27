package bootstrap

import (
	"context"
	"letschat/api/controllers"
	middlewares "letschat/api/middlewares"
	"letschat/api/repository"
	"letschat/api/routes"
	"letschat/api/services"
	infrastructure "letschat/infrastructure"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	infrastructure.Module,
	routes.Module,
	middlewares.Module,

	services.Module,
	repository.Module,

	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
	database infrastructure.Database,
	logger infrastructure.Logger,
	env infrastructure.Env,
) {
	appStop := func(context.Context) error {
		logger.Zap.Info("Stopping Application")

		// conn.Close()
		return nil
	}
	routes.Setup()
	handler.Gin.Run(":" + env.ServerPort)

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("------------------------")
			logger.Zap.Info("------  Chat Application  📺  ------")
			logger.Zap.Info("------------------------")

			go func() {
				//middlewares.Setup()
				routes.Setup()
				logger.Zap.Info("🌱 Seeding data...")
				//seeds.Run()
				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(":" + env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: appStop,
	})

}
