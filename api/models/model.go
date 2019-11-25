package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

func RetrieveLongURL(client *redis.Client, tinyURL string) (bool, string, string) {

	ok := existTinyURL(client, tinyURL)
	message := "Internal Error"
	longURL := "NONE"

	if ok {
		longURL = getURLMapping(client, tinyURL)
		message = "Success"
	}
	return ok, message, longURL
}

func CreateTinyURL(client *redis.Client, longURL string, duration time.Duration) (bool, string, string) {
	tinyURL := getFreeTinyURL(client)
	ok := false
	message := "Internal Error"

	if existTinyURL(client, tinyURL) {
		panic(errors.New("Free tinyURL from KGS already exists in DB"))
	} else {
		setURLMapping(client, tinyURL, longURL, duration)
		message = "Success"
		ok = true
	}
	return ok, message, tinyURL
}

func existTinyURL(client *redis.Client, tinyURL string) bool {
	exist, err := client.Exists(tinyURL).Result()
	if err != nil {
		panic(err)
	}
	return exist == 1
}

func getFreeTinyURL(client *redis.Client) string {
	tinyURL := "aaa"
	req := &pb.Request{A: a, B: b}
	if resp, err := addClient.Compute(ctx, req); err == nil {
		msg := fmt.Sprintf("Summation is %d", resp.Result)
		json.NewEncoder(w).Encode(msg)
	} else {
		msg := fmt.Sprintf("Internal server error: %s", err.Error())
		json.NewEncoder(w).Encode(msg)

	}
	return tinyURL
}

func setURLMapping(client *redis.Client, tinyURL string, longURL string, duration time.Duration) {
	err := client.Set(tinyURL, longURL, duration).Err()
	if err != nil {
		panic(err)
	}
}

func getURLMapping(client *redis.Client, tinyURL string) string {
	longURL, err := client.Get(tinyURL).Result()
	if err != nil {
		panic(err)
	}
	return longURL
}
