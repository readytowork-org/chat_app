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

// UserController -> struct
type UserController struct {
	logger      infrastructure.Logger
	env         infrastructure.Env
	userService services.UserService
}

// NewUserController -> constructor
func NewUserController(
	logger infrastructure.Logger,
	env infrastructure.Env,
	userService services.UserService,

) UserController {
	return UserController{
		logger:      logger,
		env:         env,
		userService: userService,
	}
}

func (cc UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		cc.logger.Zap.Error("Error creating user: [ShouldBingJSON]:", err.Error())
		err = errors.BadRequest.Wrapf(err, "JSON Binding error")
		responses.HandleError(c, err)
		return
	}

	err := cc.userService.Create(user)
	if err != nil {
		cc.logger.Zap.Error("Error creating user: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, "User Created Succesfully")
}

func (cc UserController) Update(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		msg := "Please send user id "
		cc.logger.Zap.Error(msg)
		err := errors.BadRequest.New(msg)
		responses.HandleError(c, err)
		return
	}
	var user models.UpdateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		cc.logger.Zap.Error("Error updating user: [ShouldBingJSON]:", err.Error())
		err = errors.BadRequest.Wrapf(err, "JSON Binding error")
		responses.HandleError(c, err)
		return
	}
	err := cc.userService.Update(id, user)
	if err != nil {
		cc.logger.Zap.Error("Error updating user: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, "User updated Succesfully")
}

func (cc UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		msg := "Please send user id "
		cc.logger.Zap.Error(msg)
		err := errors.BadRequest.New(msg)
		responses.HandleError(c, err)
		return
	}
	err := cc.userService.Delete(id)
	if err != nil {
		cc.logger.Zap.Error("Error deleting user: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, "User Deleted Succesfully")
}

func (cc UserController) FindOne(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		msg := "Please send user id "
		cc.logger.Zap.Error(msg)
		err := errors.BadRequest.New(msg)
		responses.HandleError(c, err)
		return
	}
	sd, err := cc.userService.FindOne(id)
	if err != nil {
		cc.logger.Zap.Error("Error finding user: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, sd)
}

func (cc UserController) FindAll(c *gin.Context) {
	sd, err := cc.userService.FindAll()
	if err != nil {
		cc.logger.Zap.Error("Error finding all user: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, sd)
}
