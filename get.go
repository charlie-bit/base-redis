package base_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// SGet set string
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

// MapGet get map
func MapGet() ([]interface{}, error) {
	err := StandAloneClient()
	if err != nil {
		return nil, err
	}

	result, err := Client.MGet(context.Background(), "test_map").Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HashGet get hash, all flied
func HashGet() (map[string]string, error) {
	err := StandAloneClient()
	if err != nil {
		return nil, err
	}

	return Client.HGetAll(context.Background(), "user_1").Result()
}

// ListGet get all list
func ListGet() ([]string, error) {
	err := StandAloneClient()
	if err != nil {
		return nil, err
	}

	return Client.LRange(context.Background(), "list_1", 0, 10).Result()
}
