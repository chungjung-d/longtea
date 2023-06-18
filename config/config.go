package config

import "os"

func getenv(key string, alterValue string) string {

	value := os.Getenv(key)

	if len(value) == 0 {
		return alterValue
	}

	return value
}
