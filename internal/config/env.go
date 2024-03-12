package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func SetupEnvVariables(envFile string) error {
	fmt.Println(envFile)
	err := godotenv.Load(envFile)
	if err != nil {
		return err
	}
	return nil
}
