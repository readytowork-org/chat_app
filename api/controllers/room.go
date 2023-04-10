package controllers

import (
	"letschat/api/services"
	"letschat/errors"
	"letschat/infrastructure"
	"letschat/models"
	"letschat/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoomController -> struct
type RoomController struct {
	logger      infrastructure.Logger
	env         infrastructure.Env
	roomService services.RoomService
}

// NewRoomController -> constructor
func NewRoomController(
	logger infrastructure.Logger,
	env infrastructure.Env,
	roomService services.RoomService,

) RoomController {
	return RoomController{
		logger:      logger,
		env:         env,
		roomService: roomService,
	}
}

func (cc RoomController) Create(c *gin.Context) {

	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		cc.logger.Zap.Error("Error creating room: [ShouldBingJSON]:", err.Error())
		err = errors.BadRequest.Wrapf(err, "JSON Binding error")
		responses.HandleError(c, err)
		return	
	}

	_, err := cc.roomService.Create(room)
	if err != nil {
		cc.logger.Zap.Error("Error creating room: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, "Room Created Succesfully")
}

func (cc RoomController) Update(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		msg := "Please send room id "
		cc.logger.Zap.Error(msg)
		err := errors.BadRequest.New(msg)
		responses.HandleError(c, err)
		return
	}
	var room models.RoomUpdate
	if err := c.ShouldBindJSON(&room); err != nil {
		cc.logger.Zap.Error("Error updating room: [ShouldBingJSON]:", err.Error())
		err = errors.BadRequest.Wrapf(err, "JSON Binding error")
		responses.HandleError(c, err)
		return
	}
	err := cc.roomService.Update(id, room)
	if err != nil {
		cc.logger.Zap.Error("Error updating room: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, "Room Updated Succesfully")
}

func (cc RoomController) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		msg := "Please send room id "
		cc.logger.Zap.Error(msg)
		err := errors.BadRequest.New(msg)
		responses.HandleError(c, err)
		return
	}
	err := cc.roomService.Delete(id)
	if err != nil {
		cc.logger.Zap.Error("Error deleting room: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, "Room Deleted Succesfully")
}

func (cc RoomController) FindOne(c *gin.Context) {

	id := c.Param("id")
	if len(id) == 0 {
		msg := "Please send room id "
		cc.logger.Zap.Error(msg)
		err := errors.BadRequest.New(msg)
		responses.HandleError(c, err)
		return
	}

	sd, err := cc.roomService.FindOne(id)
	if err != nil {
		cc.logger.Zap.Error("Error finding room: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, sd)
}
