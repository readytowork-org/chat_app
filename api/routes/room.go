package routes

import (
	"letschat/api/controllers"
	"letschat/infrastructure"
)

type RoomRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	roomController controllers.RoomController
}

func NewRoomRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	roomController controllers.RoomController,

) RoomRoutes {
	return RoomRoutes{
		router:         router,
		logger:         logger,
		roomController: roomController,
	}
}

func (c RoomRoutes) Setup() {
	c.logger.Zap.Info("Setting up user routes")

	room := c.router.Gin.Group("/room")
	{
		room.POST("", c.roomController.Create)
		room.PUT("/:id", c.roomController.Update)
		room.DELETE("/:id", c.roomController.Delete)
		room.GET("/:id", c.roomController.FindOne)

	}
}
