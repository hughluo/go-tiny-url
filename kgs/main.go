package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	u "github.com/hughluo/go-tiny-url/kgs/utils"
	"os"
	"strconv"
	"strings"
)

var CLIENT *redis.Client

func main() {
	CLIENT = CreateClient()

	//os.Setenv("REDIS_INITED", "false")
	os.Setenv("KEY_LENGTH", "2")
	if REDIS_INITED := os.Getenv("REDIS_INITED"); REDIS_INITED == "false" {
		InitRedis()
	}
	fmt.Print(GetSetFreeAmount())

}

func InitRedis() {
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
		AddToSetFree(prefix)
		return
	}

	for index := 0; index < n; index++ {
		newPrefix := prefix + charArray[index]
		addAllTinyURLToSetFreeHelper(charArray, n, newPrefix, length-1)
	}
}

func AddToSetFree(freeTinyURL string) {
	err := CLIENT.SAdd("FREE", freeTinyURL).Err()
	if err != nil {
		panic(err)
	}
}

func GetSetFreeAmount() int64 {
	amount, err := CLIENT.SCard("FREE").Result()
	if err != nil {
		panic(err)
	}
	return amount
}

func PopSetFree() string {
	freeTinyURL, err := CLIENT.SPop("FREE").Result()
	if err != nil {
		panic(err)
	}
	return freeTinyURL
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
