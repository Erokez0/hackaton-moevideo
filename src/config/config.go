package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
}

type serverConfig struct {
	Address string
	Port string
}

var DBConfig dbConfig
var ServerConfig serverConfig

func Init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	DBConfig.DbHost = os.Getenv("DB_HOST")
	DBConfig.DbUser = os.Getenv("DB_USER")
	DBConfig.DbPassword = os.Getenv("DB_PASSWORD")
	DBConfig.DbName = os.Getenv("DB_NAME")
	DBConfig.DbPort = os.Getenv("DB_PORT")

	ServerConfig.Address = os.Getenv("SERVER_ADDRESS")
	ServerConfig.Port = os.Getenv("SERVER_PORT")
}