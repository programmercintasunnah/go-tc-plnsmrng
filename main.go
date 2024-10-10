package main

import (
	"fmt"
	"go-tc-plnsmrng/api"
	"net/http"

	"github.com/go-chi/chi"
)

// Fungsi utama
func main() {
	// Inisialisasi router
	r := chi.NewRouter()
	api.SetupRoutes(r) // Setup API routes

	// Menjalankan server lokal untuk pengembangan
	httpPort := ":8080"
	fmt.Printf("Memulai server di %s\n", httpPort)
	if err := http.ListenAndServe(httpPort, r); err != nil {
		fmt.Printf("Kesalahan saat memulai server di %s: %v\n", httpPort, err)
	}
}

// Handler untuk endpoint utama
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// Fungsi utama yang diekspor untuk Vercel
func VercelHandler(w http.ResponseWriter, r *http.Request) {
	handler(w, r)
}
