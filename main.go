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
)

func main() {
	// Set up the HTTP server and routes for local development
	http.HandleFunc("/", api.MainHandler) // Use your handler from the api package
	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
