package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Printf("Starting app...")

	mapUrl()

	if err := router.Run(":8888"); err != nil {
		panic(err)
	}

}
