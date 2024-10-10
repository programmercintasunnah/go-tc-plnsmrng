package api

import (
	"encoding/json"
	"go-tc-plnsmrng/config"
	"go-tc-plnsmrng/internal/handlers"
	"go-tc-plnsmrng/internal/repository"
	"net/http"
	"os"

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
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Define routes
	router.Post("/api/bobot", bobotHandler.CreateBobot)
	router.Get("/api/bobots", bobotHandler.GetAllBobots)
	router.Get("/api/jsondata", NewJSONHandler())
}

// MainHandler is the entry point for AWS Lambda
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// This function can remain empty or be used for Lambda specific handling
}

// NewJSONHandler returns a handler for serving JSON data
func NewJSONHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("data.json") // Open the JSON file
		if err != nil {
			http.Error(w, "Could not open data file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		var data []map[string]interface{} // Change to your desired structure
		if err := json.NewDecoder(file).Decode(&data); err != nil {
			http.Error(w, "Could not decode data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data) // Encode data as JSON
	}
}
