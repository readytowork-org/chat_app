package routes

import (
	"letschat/api/controllers"
	"letschat/infrastructure"
)

type MessageRoutes struct {
	logger            infrastructure.Logger
	router            infrastructure.Router
	messageController controllers.MessageController
}

func NewMessageRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	messageController controllers.MessageController,

) MessageRoutes {
	return MessageRoutes{
		router:            router,
		logger:            logger,
		messageController: messageController,
	}
}

func (c MessageRoutes) Setup() {
	c.logger.Zap.Info("Setting up user routes")

	message := c.router.Gin.Group("/messages")
	{
		message.DELETE("/:id", c.messageController.Delete)
		message.GET("/:roomid", c.messageController.FindAll)
	}
}
