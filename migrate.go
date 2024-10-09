package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

//cara menjalankannya
// go run migrate.go up
// go run migrate.go down

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	m, err := migrate.New(
		"file://./migrations",
		dbURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Menjalankan migrasi
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "up":
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		case "down":
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		default:
			log.Println("Unknown command")
		}
	}
}
