package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DB        *sqlx.DB
	JWTSecret string
}

func NewConfig() *Config {
	db, err := sqlx.Connect("postgres", "user=youruser password=yourpassword dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return &Config{
		DB:        db,
		JWTSecret: "supersecretkey",
	}
}
