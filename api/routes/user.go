package routes

import (
	"letschat/api/controllers"
	"letschat/infrastructure"
)

type UserRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	userController controllers.UserController
}

func NewUserRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	userController controllers.UserController,

) UserRoutes {
	return UserRoutes{
		router:         router,
		logger:         logger,
		userController: userController,
	}
}

func (c UserRoutes) Setup() {
	c.logger.Zap.Info("Setting up user routes")

	user := c.router.Gin.Group("/user")
	{
		user.POST("", c.userController.Create)
		user.PUT("/:id", c.userController.Update)
		user.DELETE("/:id", c.userController.Delete)
		user.GET("/:id", c.userController.FindOne)
		user.GET("", c.userController.FindAll)
	}
}
