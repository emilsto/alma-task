package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	MongoURI        string
	MongoCollection string
}

// Returns a Config struct with the values from the .env file
func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Port:            os.Getenv("PORT"),
		MongoURI:        os.Getenv("MONGO_URI"),
		MongoCollection: os.Getenv("MONGO_COLLECTION"),
	}
}
