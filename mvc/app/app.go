package app

import (
	"log"
	"net/http"

	"github.com/mkruczek/go-service-first/mvc/controllers"
)

func StartApp() {
	log.Printf("Starting app...")
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/user", controllers.GetUser)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}

}
