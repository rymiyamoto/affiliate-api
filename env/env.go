package env

import (
	"fmt"
	"os"
)

func GetValue(key string) (string, error) {
	v := os.Getenv("IS_DEV_APP")

	if v == "" {
		return "", fmt.Errorf("not found value. key: %s", key)
	}

	return v, nil
}
