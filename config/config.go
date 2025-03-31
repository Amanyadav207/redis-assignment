package config

import (
	"os"
)

const DefaultPort = "7171"

func GetPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return DefaultPort
}