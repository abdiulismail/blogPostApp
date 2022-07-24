package config

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
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
