package api

import (
	"go-tc-plnsmrng/config"
	"go-tc-plnsmrng/internal/handlers"
	"go-tc-plnsmrng/internal/repository"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// SetupRoutes sets up the API routes
func SetupRoutes(router *chi.Mux) {
	// Initialize the application configuration
	cfg := config.NewConfig()
	repo := repository.NewBobotRepository(cfg.DB)
	bobotHandler := handlers.NewBobotHandler(repo)

	// Middleware CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Define routes
	router.Post("/api/bobot", bobotHandler.CreateBobot)
	router.Get("/api/bobots", bobotHandler.GetAllBobots)
}

// MainHandler is the entry point for AWS Lambda
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// This function can remain empty or be used for Lambda specific handling
}
