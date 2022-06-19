package base_redis

import (
	"context"
	"time"
)

// SSet set string
func SSet() (string, error) {
	err := StandAloneClient()
	if err != nil {
		return "", err
	}

	result, err := Client.Set(context.Background(), "test", "test", time.Minute*1).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

// MapSet set map
func MapSet() (string, error) {
	err := StandAloneClient()
	if err != nil {
		return "", err
	}

	return Client.MSet(context.Background(), map[string]interface{}{"key1": "value1", "key2": "value2"}).Result()
}

// HashSet set hash
func HashSet() (int64, error) {
	err := StandAloneClient()
	if err != nil {
		return 0, err
	}

	return Client.HSet(context.Background(), "user_1", map[string]interface{}{"name": "charlie", "age": 24}).Result()
}

func ListSet() (int64, error) {
	err := StandAloneClient()
	if err != nil {
		return 0, err
	}

	return Client.LPush(context.Background(), "list_1", "1", "2", "3").Result()
}
