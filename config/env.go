package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost, Port, DBUser, DBPassword, DBAddress, DBName string
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Warning: Could not load .env file: %v", err)
	}
	dbHost := getEnv("DB_HOST")
	dbPort := getEnv("DB_PORT")

	return Config{
		PublicHost: getEnv("PUBLIC_HOST"),
		Port:       getEnv("PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBAddress:  fmt.Sprintf("%s:%s", dbHost, dbPort),
		DBName:     getEnv("DB_NAME"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key); 
	if !ok {
		log.Fatalf("Environment variable %s is required but not set", key)
	}

	return value
}