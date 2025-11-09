package config

import (
	"os"
	"strconv"
)

func getString(key, defaulValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaulValue
	}
	return value
}

func getInt(key string, defaulValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)

	if err != nil {
		return defaulValue
	}

	return i
}

func getBool(key string, defaulValue bool) bool {
	value := os.Getenv(key)
	b, err := strconv.ParseBool(value)

	if err != nil {
		return defaulValue
	}

	return b
}
