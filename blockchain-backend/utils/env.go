// utils/env.go
package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return err
	}
	return nil
}
