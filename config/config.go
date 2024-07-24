package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type setting struct {
	Host     string
	Port     int
	Name     string
	Email    string
	Password string
}

func ImportSetting() setting {
	var result setting

	err := godotenv.Load(".env")

	if err != nil {
		return setting{}
	}

	result.Host = os.Getenv("host")
	result.Port, _ = strconv.Atoi(os.Getenv("port"))
	result.Name = os.Getenv("name")
	result.Email = os.Getenv("email")
	result.Password = os.Getenv("password")
	return result
}
