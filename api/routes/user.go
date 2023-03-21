package routes

import (
	"letschat/api/controllers"
	"letschat/api/middlewares"
	"letschat/infrastructure"
)

type UserRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	userController controllers.UserController
	jwtMiddleware  middlewares.JWTAuthMiddleWare
}

func NewUserRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	userController controllers.UserController,
	jwtMiddleware middlewares.JWTAuthMiddleWare,

) UserRoutes {
	return UserRoutes{
		router:         router,
		logger:         logger,
		userController: userController,
		jwtMiddleware:  jwtMiddleware,
	}
}

func (c UserRoutes) Setup() {
	c.logger.Zap.Info("Setting up user routes")

	user := c.router.Gin.Group("/user")
	{
		user.POST("", c.userController.Create)
	}
	authUser := user.Use(c.jwtMiddleware.Handle())
	{
		authUser.PUT("/:id", c.userController.Update)
		authUser.DELETE("/:id", c.userController.Delete)
		authUser.GET("/:id", c.userController.FindOne)
		authUser.GET("", c.userController.FindAll)
	}
}
