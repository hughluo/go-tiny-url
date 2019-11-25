package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/hughluo/go-tiny-url/api/controllers"
	"github.com/hughluo/go-tiny-url/api/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"

	"github.com/hughluo/go-tiny-url/pb"
	"google.golang.org/grpc"
)

func main() {
	// Configure redis client
	redisClient := CreateClient()
	controllers.SetRedisClient(redisClient)

	// Configure gRPC client
	GRPC_DIAL_TARGET := os.Getenv("GRPC_DIAL_TARGET")
	conn, err := grpc.Dial(GRPC_DIAL_TARGET, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	kgsClient := pb.NewKGSServiceClient(conn)
	models.SetKGSClient(kgsClient)

	// Configure REST API router
	router := httprouter.New()
	router.POST("/gotinyurl/", controllers.CreateTinyURL)
	router.GET("/gotinyurl/:tinyurl", controllers.RetrieveLongURL)

	API_ADDRESS := os.Getenv("API_ADDRESS")
	log.Fatal(http.ListenAndServe(API_ADDRESS, router))

}

func CreateClient() *redis.Client {
	REDIS_MAIN_ADDRESS := os.Getenv("REDIS_MAIN_ADDRESS")
	REDIS_MAIN_PASSWORD := os.Getenv("REDIS_MAIN_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_MAIN_ADDRESS,
		Password: REDIS_MAIN_PASSWORD,
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}
