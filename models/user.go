package models

import (
	"blog/config"
	"errors"
	"github.com/go-redis/redis/v9"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidLogin = errors.New("invalid user")
)

func RegisterUser(username, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	return Client.Set(config.Ctx, "user:"+username, hash, 0).Err()
}

func AuthenticateUser(username, password string) error {
	hash, err := Client.Get(config.Ctx, "user:"+username).Bytes()
	if err == redis.Nil {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return ErrInvalidLogin
	}
	return err
}
