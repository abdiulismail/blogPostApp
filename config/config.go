package config

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var Mysession = sessions.NewCookieStore([]byte("dhfkshgdfj"))

var Templates *template.Template

var Port = ":8080"

var Client *redis.Client

var Ctx = context.TODO()

func CheckError(e error) {
	if e != nil {
		log.Println(e.Error())
	}
}

func AuthRequired(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Mysession.Get(r, "session")
		_, ok := session.Values["username"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		h.ServeHTTP(w, r)
	}
}
