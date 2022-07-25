package main

import (
	"blog/config"
	"blog/models"
	"blog/routes"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {

	models.Init()

	//parse all template filesx
	config.Templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()

	r.HandleFunc("/login", routes.LoginGetHandler).Methods("GET")
	r.HandleFunc("/login", routes.LoginPostHandler).Methods("POST")
	r.HandleFunc("/logout", routes.LogoutPostHandler).Methods("GET")

	r.HandleFunc("/register", routes.RegisterGetHandler).Methods("GET")
	r.HandleFunc("/register", routes.RegisterPostHandler).Methods("POST")

	r.HandleFunc("/", config.AuthRequired(routes.IndexGetHandler)).Methods("GET")
	r.HandleFunc("/", config.AuthRequired(routes.IndexPostHandler)).Methods("POST")

	//serve static files
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)

	log.Println("starting application on port 8080")
	_ = http.ListenAndServe(config.Port, r)
}
