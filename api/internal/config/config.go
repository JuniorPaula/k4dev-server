package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Port = 0

func InitEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}
}
