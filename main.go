package main

import (
	"go-tc-plnsmrng/config"
	"go-tc-plnsmrng/internal/handlers"
	"go-tc-plnsmrng/internal/repository"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()
	repo := repository.NewBobotRepository(cfg.DB)
	bobotHandler := handlers.NewBobotHandler(repo)

	router := chi.NewRouter()

	// Middleware CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Post("/api/bobot", bobotHandler.CreateBobot)
	router.Get("/api/bobots", bobotHandler.GetAllBobots)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	router.ServeHTTP(w, r) // Menjalankan router sebagai HTTP handler
}
