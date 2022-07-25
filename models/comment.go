package models

import "blog/config"

func GetComments() ([]string, error) {

	return Client.LRange(config.Ctx, "comments", 0, 100).Result()
}

func PostComments(comment string) error {
	return Client.LPush(config.Ctx, "comments", comment).Err()
}
