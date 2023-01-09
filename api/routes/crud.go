package routes

import (
	controller "letschat/api/controllers"
	"letschat/api/middlewares"
	"letschat/infrastructure"
)

type CrudRoutes struct {
	crudController controller.CrudController
	router         infrastructure.Router
	middleware     middlewares.FirebaseAuth
}

func NewCrudRoutes(
	crudController controller.CrudController,
	router infrastructure.Router,
	middleware middlewares.FirebaseAuth,
) CrudRoutes {
	return CrudRoutes{
		crudController: crudController,
		router:         router,
		middleware:     middleware,
	}
}

func (cr CrudRoutes) Setup() {

	crud := cr.router.Gin.Group("firebase").Use(cr.middleware.Handle())
	{

		crud.POST("", cr.crudController.CreateData)
		crud.GET("", cr.crudController.GetData)
		crud.DELETE("", cr.crudController.UpdateData)
		crud.PUT("", cr.crudController.UpdateData)
	}
}
