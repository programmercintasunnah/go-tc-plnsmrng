package handlers

import (
	"go-tc-plnsmrng/config"
	"go-tc-plnsmrng/internal/repository"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Konfigurasi aplikasi
	cfg := config.NewConfig()
	repo := repository.NewBobotRepository(cfg.DB)
	bobotHandler := NewBobotHandler(repo)

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

	router.ServeHTTP(w, r)
}
