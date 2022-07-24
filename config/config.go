package config

import (
	"context"
	"github.com/go-redis/redis/v9"
	"html/template"
	"log"
)

var Templates *template.Template

var Port = ":8080"

var Client *redis.Client

var Ctx = context.TODO()

func CheckError(e error) {
	if e != nil {
		log.Println(e.Error())
	}
}
