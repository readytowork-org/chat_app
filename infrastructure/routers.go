package infrastructure

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin *gin.Engine
}

func NewRouter(fbauth *auth.Client) Router {
	httpRouter := gin.Default()
	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api is runnings."})
	})
	httpRouter.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", fbauth)
	})
	return Router{
		Gin: httpRouter,
	}
}
