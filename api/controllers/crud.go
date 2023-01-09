package controllers

import (
	"context"
	"fmt"
	"letschat/api/services"
	"letschat/infrastructure"
	"letschat/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CrudController struct {
	logger   infrastructure.Logger
	services services.CrudService
	db       infrastructure.Database
}

func NewCrudController(
	logger infrastructure.Logger,
	services services.CrudService,
	db infrastructure.Database,
) CrudController {
	return CrudController{
		logger:   logger,
		services: services,
		db:       db,
	}

}

func (cc CrudController) CreateData(ctx *gin.Context) {
	uUID := ctx.MustGet("UUID")
	ref := cc.db.DB.NewRef("details/" + fmt.Sprint(uUID))
	var details models.UserDetails
	if err := ctx.ShouldBindJSON(&details); err != nil {
		msg := "Error validating details input"
		cc.logger.Zap.Info(msg, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": msg,
		})
		return
	}

	if err := ref.Set(context.TODO(), details); err != nil {
		log.Fatal(err)
	}

	fmt.Println("score added/updated successfully!")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
	return
}

func (cc CrudController) UpdateData(ctx *gin.Context) {}
func (cc CrudController) DeleteData(ctx *gin.Context) {}
func (cc CrudController) GetData(ctx *gin.Context)    {}
