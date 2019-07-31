package core

import "github.com/go-redis/redis"

var Redis *redis.Client

func InitRedis() error {
	c, err := Conf.Map("redis")
	if err != nil {
		return err
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     c["Addr"].(string),
		Password: c["Password"].(string),
		DB:       c["DB"].(int),
	})
	return nil
}
