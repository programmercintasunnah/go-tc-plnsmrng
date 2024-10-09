package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DB        *sqlx.DB
	JWTSecret string
}

func NewConfig() *Config {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the database URL from the .env file
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalln("DB_URL not found in environment")
	}

	// Connect to the database using the URL from .env
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}

	return &Config{
		DB:        db,
		JWTSecret: "supersecretkey", // You might want to move this to .env as well
	}
}
