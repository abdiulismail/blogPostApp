package main

import (
	"blog/config"
	"blog/handlers"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {

	config.Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	//parse all template filesx
	config.Templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.LoginGetHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginPostHandler).Methods("POST")

	r.HandleFunc("/register", handlers.RegisterGetHandler).Methods("GET")
	r.HandleFunc("/register", handlers.RegisterPostHandler).Methods("POST")

	r.HandleFunc("/", handlers.IndexGetHandler).Methods("GET")
	r.HandleFunc("/", handlers.IndexPostHandler).Methods("POST")

	//serve static files
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)

	log.Println("starting application on port 8080")
	_ = http.ListenAndServe(config.Port, r)
}
