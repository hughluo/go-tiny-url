package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateTinyURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprintf(w, "create \n")
}

func RetrieveTinyURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "retrieve, %s!\n", ps.ByName("tinyurl"))
}

func UpdateTinyURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "update, %s!\n", ps.ByName("tinyurl"))
	fmt.Fprintf(w, "Not implemented!")
}

func DeleteTinyURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "delete, %s!\n", ps.ByName("tinyurl"))
	fmt.Fprintf(w, "Not implemented!")
}
