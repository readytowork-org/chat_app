package routes

import (
	controller "letschat/api/controllers"
	"letschat/infrastructure"
)

type FirebaseRoutes struct {
	firebaseController controller.FirebaseController
	router             infrastructure.Router
}

func NewFirebaseRoutes(firebaseController controller.FirebaseController, router infrastructure.Router) FirebaseRoutes {
	return FirebaseRoutes{
		firebaseController: firebaseController,
		router:             router,
	}
}

func (fr FirebaseRoutes) Setup() {
	firebase := fr.router.Gin.Group("user").Use()
	{
		firebase.POST("/", fr.firebaseController.CreateUser)
		firebase.POST("/login", fr.firebaseController.LoginUser)
	}
}
