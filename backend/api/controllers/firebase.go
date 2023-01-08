package controllers

import (
	"letschat/api/services"
	"letschat/infrastructure"
	"letschat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FirebaseController struct {
	logger   infrastructure.Logger
	services services.FirebaseService
}

func NewFirebaseController(
	logger infrastructure.Logger,
	services services.FirebaseService,
) FirebaseController {
	return FirebaseController{
		logger:   logger,
		services: services,
	}
}

func (fc FirebaseController) CreateUser(ctx *gin.Context) {
	var newUser models.FirebaseAuthUser
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		msg := "Error validating user input"
		fc.logger.Zap.Info(msg, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": msg,
		})
		return
	}
	registeredUser, err := fc.services.CreateUser(newUser)
	if err != nil {

		msg := "Error validating user input"
		fc.logger.Zap.Info(msg, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": msg,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"data":    registeredUser,
	})

}

func (fc FirebaseController) LoginUser(ctx *gin.Context) {
	var newUser models.FirebaseAuthUser
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		msg := "Error validating user input"
		fc.logger.Zap.Info(msg, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": msg,
		})
		return
	}
	token, err := fc.services.LoginUser(newUser)
	if err != nil {

		msg := "Error validating user input"
		fc.logger.Zap.Info(msg, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": msg,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Logged In",
		"data":    token,
	})

}
