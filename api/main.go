package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/hughluo/go-tiny-url/api/controllers"
	"github.com/hughluo/go-tiny-url/api/models"
	"github.com/julienschmidt/httprouter"

	"github.com/hughluo/go-tiny-url/pb"
	"google.golang.org/grpc"
)

func main() {
	// Configure redis client
	redisClient := CreateClient()
	controllers.SetRedisClient(redisClient)

	// Configure gRPC client
	conn, err := grpc.Dial("kgs-service:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	kgsClient := pb.NewKGSServiceClient(conn)
	models.SetKGSClient(kgsClient)

	// Configure REST API router
	router := httprouter.New()
	router.POST("/gotinyurl/", controllers.CreateTinyURL)
	router.GET("/gotinyurl/:tinyurl", controllers.RetrieveLongURL)
	// router.PUT("/gotinyurl/:tinyurl", controllers.UpdateTinyURL)
	// router.DELETE("/gotinyurl/:tinyurl", controllers.DeleteTinyURL)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func CreateClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:7002",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}
