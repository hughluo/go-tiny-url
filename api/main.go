package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/hughluo/go-tiny-url/api/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	client := CreateClient()
	controllers.SetClient(client)
	router := httprouter.New()
	router.POST("/gotinyurl/", controllers.CreateTinyURL)
	router.GET("/gotinyurl/:tinyurl", controllers.RetrieveTinyURL)
	// router.PUT("/gotinyurl/:tinyurl", controllers.UpdateTinyURL)
	// router.DELETE("/gotinyurl/:tinyurl", controllers.DeleteTinyURL)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func CreateClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:7001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}
