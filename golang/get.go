package golang

/**
set string
*/
func SGet() (string, error) {
	err := StandAloneClient()
	if err != nil {
		return "", err
	}

	result, err := Client.Get("test").Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
