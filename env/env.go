package env

import (
	"fmt"
	"os"
)

func GetValue(key string) (string, error) {
	v := os.Getenv("key")

	if v == "" {
		return "", fmt.Errorf("not found value. key: %s", key)
	}

	return v, nil
}
