package base_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

/**
set string
*/
func SGet() (string, error) {
	err := StandAloneClient()
	if err != nil {
		return "", err
	}

	result, err := Client.Get(context.Background(), "test").Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}

	return result, nil
}
