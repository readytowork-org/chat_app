package bootstrap

import (
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
	handler infrastructure.Router,
	routes routes.Routes,
	env infrastructure.Env,
) {
	routes.Setup()
	handler.Gin.Run(":" + env.ServerPort)

}
