package lib

import (
	"fmt"
	"os"
)

func checkEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Println("Error: Environment variable", key, "is not set.")
		os.Exit(1)
	}
	return value
}
