package main

import (
	"blog/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/hello",handlers.MyHelloHandler).Methods("GET")
	r.HandleFunc("/goodbye",handlers.GoodBye).Methods("GET")


	log.Println("starting application on port 8080")
	http.ListenAndServe(":8080",r)
}
