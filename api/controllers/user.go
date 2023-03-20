package controllers

import (
	"letschat/api/services"
	"letschat/api/validators"
	"letschat/errors"
	"letschat/infrastructure"
	"letschat/models"
	"letschat/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController -> struct
type UserController struct {
	logger        infrastructure.Logger
	env           infrastructure.Env
	userService   services.UserService
	userValidator validators.UserValidator
}

// NewUserController -> constructor
func NewUserController(
	logger infrastructure.Logger,
	env infrastructure.Env,
	userService services.UserService,
	userValidator validators.UserValidator,

) UserController {
	return UserController{
		logger:        logger,
		env:           env,
		userService:   userService,
		userValidator: userValidator,
	}
}

func (cc UserController) Create(c *gin.Context) {
	var user models.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		cc.logger.Zap.Error("Error creating user: [ShouldBingJSON]:", err.Error())
		err = errors.BadRequest.Wrapf(err, "JSON Binding error")
		responses.HandleError(c, err)
		return
	}
	if validationErr := cc.userValidator.Validate.Struct(user); validationErr != nil {
		err := errors.BadRequest.Wrap(validationErr, "Validation error")
		err = errors.SetCustomMessage(err, "Invalid input information")
		err = errors.AddErrorContextBlock(err, cc.userValidator.GenerateValidationResponse(validationErr))
		responses.HandleError(c, err)
		return
	}
	if user.Password != user.ConfirmPassword {
		responses.ErrorJSON(c, http.StatusBadRequest, "Password and Confirm Password doesn't match")
		return
	}
	_, exits, err := cc.userService.CheckUserWithPhone(user.PhoneNumber)
	if err != nil {
		cc.logger.Zap.Error("Error while creating user: ", err.Error())
		responses.HandleError(c, err)
		return
	}
	if exits {
		responses.ErrorJSON(c, http.StatusBadRequest, "User already exists")
		return
	}
	err = user.BeforeCreate()
	if err != nil {
		cc.logger.Zap.Error("Error while encryptinh password: ", err.Error())
		responses.HandleError(c, err)
	}
	err = cc.userService.Create(user)
	if err != nil {
		cc.logger.Zap.Error("Error while creating user: ", err.Error())
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
