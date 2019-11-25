package controllers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/hughluo/go-tiny-url/models"
	"github.com/julienschmidt/httprouter"
	"github.com/go-redis/redis/v7"

)

var CLIENT *redis.Client

func SetClient(client *redis.Client) {
	CLIENT = client
}

func CreateTinyURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	longURL := r.Form["longURL"][0]
	ok, message, tinyURL := models.CreateTinyURL(CLIENT, longURL)
	if !ok {
		log.Printf("create tiny url failed! %s", message)
		w.WriteHeader(400)
	}

	fmt.Fprintf(w, "tinyurl created %s!", tinyURL)
}

func RetrieveTinyURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ok, message, longURL := models.RetrieveLongURL(CLIENT, tinyURL)
}

func UpdateTinyURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "update, %s!\n", ps.ByName("tinyurl"))
	fmt.Fprintf(w, "Not implemented!")
}

func DeleteTinyURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "delete, %s!\n", ps.ByName("tinyurl"))
	fmt.Fprintf(w, "Not implemented!")
}
