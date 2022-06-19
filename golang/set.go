package golang

import "time"

// set string
func SSet() (string, error) {
	err := StandAloneClient()
	if err != nil {
		return "", err
	}

	result, err := Client.Set("test", "test", time.Minute*1).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
