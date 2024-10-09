package main

import (
	"go-tc-plnsmrng/config"
	"go-tc-plnsmrng/internal/handlers"
	"go-tc-plnsmrng/internal/repository"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.NewConfig()
	repo := repository.NewBobotRepository(cfg.DB)
	bobotHandler := handlers.NewBobotHandler(repo)

	r := chi.NewRouter()

	r.Post("/bobot", bobotHandler.CreateBobot)
	r.Get("/bobots", bobotHandler.GetAllBobots)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
