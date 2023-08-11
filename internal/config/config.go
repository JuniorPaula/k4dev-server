package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectStringMySQL  = ""
	ConnetStringMongoDB = ""
	FrontendURL         = ""
	Port                = 0
	SecretKey           []byte
)

func InitEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}

	ConnectStringMySQL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)

	if os.Getenv("NODE_ENV") == "production" {
		ConnetStringMongoDB = os.Getenv("MONGO_URL")
		return
	} else if os.Getenv("NODE_ENV") == "development" {
		ConnetStringMongoDB = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
			os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
			os.Getenv("MONGO_HOST"),
			os.Getenv("MONGO_PORT"),
		)
		return
	}

	SecretKey = []byte(os.Getenv("JWT_SECRET"))

	FrontendURL = os.Getenv("FRONTEND_URL")
}
