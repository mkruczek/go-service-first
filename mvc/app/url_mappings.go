package app

import (
	"github.com/mkruczek/go-service-first/mvc/controllers"
)

func mapUrl() {

	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
}
