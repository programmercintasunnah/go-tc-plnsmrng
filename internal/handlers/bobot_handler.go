package handlers

import (
	"encoding/json"
	"go-tc-plnsmrng/internal/models"
	"go-tc-plnsmrng/internal/repository"
	"net/http"
	"strings"
)

type BobotHandler struct {
	repo *repository.BobotRepository
}

func NewBobotHandler(repo *repository.BobotRepository) *BobotHandler {
	return &BobotHandler{repo: repo}
}

func (h *BobotHandler) CreateBobot(w http.ResponseWriter, r *http.Request) {
	var bobotSpec models.BobotSpec
	if err := json.NewDecoder(r.Body).Decode(&bobotSpec); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validasi nomor yang diinput
	if bobotSpec.Nomor == "" || bobotSpec.Nomor[len(bobotSpec.Nomor)-1] == '.' {
		http.Error(w, "Nomor tidak boleh kosong atau diakhiri dengan titik.", http.StatusBadRequest)
		return
	}

	// Cek apakah nomor sudah ada di database
	existingBobot, err := h.repo.GetBobotByNomor(bobotSpec.Nomor)
	if err != nil {
		http.Error(w, "Gagal melakukan pengecekan pada nomor.", http.StatusInternalServerError)
		return
	}
	if existingBobot != nil {
		http.Error(w, "Nomor sudah ada.", http.StatusBadRequest)
		return
	}

	// Buat instance Bobot
	var bobot models.Bobot
	bobot.Nama = bobotSpec.Nama
	bobot.Nomor = bobotSpec.Nomor

	// Pisahkan nomor berdasarkan titik
	parts := strings.Split(bobotSpec.Nomor, ".")

	// Logika untuk menentukan ParentID
	if len(parts) == 1 {
		// Jika nomor seperti "1", "2", "3", parent_id harus null
		bobot.ParentID = nil
	} else {
		// Gabungkan parts kecuali yang terakhir untuk mencari parent
		parentNomor := strings.Join(parts[:len(parts)-1], ".")

		// Cek apakah parent ada di database
		parentBobot, err := h.repo.GetBobotByNomor(parentNomor)
		if err != nil {
			http.Error(w, "Gagal melakukan pengecekan pada parent nomor.", http.StatusInternalServerError)
			return
		}
		if parentBobot == nil {
			http.Error(w, "Parent nomor tidak ditemukan.", http.StatusBadRequest)
			return
		}
		bobot.ParentID = &parentBobot.ID
	}

	// Simpan bobot ke database
	if err := h.repo.CreateBobot(&bobot); err != nil {
		http.Error(w, "Gagal menyimpan bobot.", http.StatusInternalServerError)
		return
	}

	// Buat response
	bobotResponse := models.BobotResponse{
		Nama:  bobot.Nama,
		Nomor: bobot.Nomor,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bobotResponse)
}



func (h *BobotHandler) GetAllBobots(w http.ResponseWriter, r *http.Request) {
	bobots, err := h.repo.GetAllBobots()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengonversi hasil query menjadi BobotResponse
	var bobotResponses []models.BobotResponse
	for _, bobot := range bobots {
		bobotResponse := models.BobotResponse{
			Nama:  bobot.Nama,
			Nomor: bobot.Nomor,
		}
		bobotResponses = append(bobotResponses, bobotResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bobotResponses)
}
