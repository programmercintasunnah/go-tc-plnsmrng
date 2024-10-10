// package main

// import (
// 	"go-tc-plnsmrng/config"
// 	"go-tc-plnsmrng/internal/handlers"
// 	"go-tc-plnsmrng/internal/repository"
// 	"log"
// 	"net/http"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/cors"
// )

// // Struktur Proyek:
// // /bobot-service
// //   ├── /main.go			   # Entry point
// //   ├── /internal
// //   │     ├── /handlers       # HTTP handlers
// //   │     ├── /models         # Model struct
// //   │     └── /repository     # DB interactions
// //   ├── /config
// //   │     └── /config.go      # Database and JWT config
// //   └── /migrations           # DB migration scripts

// func main() {
// 	cfg := config.NewConfig()
// 	repo := repository.NewBobotRepository(cfg.DB)
// 	bobotHandler := handlers.NewBobotHandler(repo)

// 	r := chi.NewRouter()

// 	// Middleware CORS
// 	r.Use(cors.Handler(cors.Options{
// 		AllowedOrigins:   []string{"*"},                             // Allow all origins (set as needed)
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // HTTP methods that are allowed
// 		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allowed headers
// 		ExposedHeaders:   []string{"Link"},
// 		AllowCredentials: false,
// 		MaxAge:           300, // Maximum value for preflight request caching
// 	}))

// 	r.Post("/api/bobot", bobotHandler.CreateBobot)
// 	r.Get("/api/bobots", bobotHandler.GetAllBobots)

// 	log.Println("Starting server on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// // Handler untuk endpoint utama
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, World!")
// }

// // Fungsi utama yang diekspor untuk Vercel
// func VercelHandler(w http.ResponseWriter, r *http.Request) {
// 	handler(w, r)
// }

package main

import (
	"fmt"
	"go-tc-plnsmrng/api"
	"net/http"
	"os"

	"github.com/akrylysov/algnhsa"
	"github.com/go-chi/chi"
)

func main() {
	// Initialize the router
	r := chi.NewRouter()
	api.SetupRoutes(r) // Set up the API routes

	// Check if running in AWS Lambda environment
	if _, ok := os.LookupEnv("_LAMBDA_SERVER_PORT"); ok {
		// Use algnhsa to listen and serve for AWS Lambda
		algnhsa.ListenAndServe(r, nil)
	} else {
		// Run a local HTTP server for development
		httpPort := ":3000"
		fmt.Printf("Starting server on %s\n", httpPort)
		if err := http.ListenAndServe(httpPort, r); err != nil {
			fmt.Printf("Error starting server on %s: %v\n", httpPort, err)
		}
	}
}
