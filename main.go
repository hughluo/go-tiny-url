package main

import (
	"log"
	"net/http"

	"github.com/hughluo/go-tiny-url/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/gotinyurl/", controllers.CreateTinyURL)
	router.GET("/gotinyurl/:tinyurl", controllers.RetrieveTinyURL)
	// router.PUT("/gotinyurl/:tinyurl", controllers.UpdateTinyURL)
	// router.DELETE("/gotinyurl/:tinyurl", controllers.DeleteTinyURL)

	log.Fatal(http.ListenAndServe(":8080", router))
}
