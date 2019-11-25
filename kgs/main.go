package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v7"
	u "github.com/hughluo/go-tiny-url/kgs/utils"

	"github.com/hughluo/go-tiny-url/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type gRPCServer struct{}

var CLIENT *redis.Client

func main() {
	// set up redis client
	CLIENT = createClient()

	//os.Setenv("REDIS_INITED", "false")
	os.Setenv("KEY_LENGTH", "2")
	if REDIS_INITED := os.Getenv("REDIS_INITED"); REDIS_INITED == "false" {
		initRedis()
	}
	fmt.Print(getSetFreeAmount())

	// Set up gRPC
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKGSServiceServer(s, &gRPCServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *gRPCServer) GetFreeGoTinyURL(cxt context.Context, req *pb.KGSRequest) (*pb.KGSResponse, error) {
	result := &pb.KGSResponse{}
	result.Result = popSetFree()
	logMessage := fmt.Sprintf("KGS req: %s result: %s", req, result.Result)
	log.Println(logMessage)
	return result, nil
}

func initRedis() {
	KEY_LENGTH, err := strconv.Atoi(os.Getenv("KEY_LENGTH"))
	if err != nil {
		panic(err)
	}
	base62 := u.GetBase62String()
	base62Slice := strings.Split(base62, "")
	addAllTinyURLToSetFree(base62Slice, KEY_LENGTH)
}

func addAllTinyURLToSetFree(charArray []string, keyLength int) {
	addAllTinyURLToSetFreeHelper(charArray, len(charArray), "", keyLength)
}

func addAllTinyURLToSetFreeHelper(charArray []string, n int, prefix string, length int) {
	if length == 0 {
		addToSetFree(prefix)
		return
	}

	for index := 0; index < n; index++ {
		newPrefix := prefix + charArray[index]
		addAllTinyURLToSetFreeHelper(charArray, n, newPrefix, length-1)
	}
}

func addToSetFree(freeTinyURL string) {
	err := CLIENT.SAdd("FREE", freeTinyURL).Err()
	if err != nil {
		panic(err)
	}
}

func getSetFreeAmount() int64 {
	amount, err := CLIENT.SCard("FREE").Result()
	if err != nil {
		panic(err)
	}
	return amount
}

func popSetFree() string {
	freeTinyURL, err := CLIENT.SPop("FREE").Result()
	if err != nil {
		panic(err)
	}
	return freeTinyURL
}

func createClient() *redis.Client {
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
