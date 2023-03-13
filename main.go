package main

import (
	"log"
	"net/http"

	"github.com/fadyboy/micro1/api"
	"github.com/gorilla/mux"
)

func main() {
	l := log.Default()

	a := api.NewAPI(l)

	r := mux.NewRouter()

	r.HandleFunc("/", a.Hello)
	r.HandleFunc("/health", a.Health)

	log.Println("Starting server on port 80")
	log.Fatal(http.ListenAndServe(":80", r))
}
