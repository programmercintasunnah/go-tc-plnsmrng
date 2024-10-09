package handlers

import (
	"encoding/json"
	"go-tc-plnsmrng/internal/models"
	"go-tc-plnsmrng/internal/repository"
	"net/http"
	"regexp"
	"strings"
)

type BobotHandler struct {
	repo *repository.BobotRepository
}

func NewBobotHandler(repo *repository.BobotRepository) *BobotHandler {
	return &BobotHandler{repo: repo}
}

// Struct untuk respons standar
type APIResponse struct {
	Status string      `json:"status"`
	Msg    string      `json:"message,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (h *BobotHandler) CreateBobot(w http.ResponseWriter, r *http.Request) {
	var bobotSpec models.BobotSpec
	if err := json.NewDecoder(r.Body).Decode(&bobotSpec); err != nil {
		response := APIResponse{Status: "error", Msg: "Invalid input"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validasi nomor yang diinput
	if bobotSpec.Nomor == "" || bobotSpec.Nomor[len(bobotSpec.Nomor)-1] == '.' {
		response := APIResponse{Status: "error", Msg: "Nomor tidak boleh kosong atau diakhiri dengan titik."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Cek apakah nomor hanya terdiri dari angka dan titik
	validNomor := regexp.MustCompile(`^(\d+(\.\d+)*)?$`)
	if !validNomor.MatchString(bobotSpec.Nomor) {
		response := APIResponse{Status: "error", Msg: "Nomor harus berupa angka yang valid (contoh: 1, 1.1, 2.3)"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Cek apakah nomor sudah ada di database
	existingBobot, err := h.repo.GetBobotByNomor(bobotSpec.Nomor)
	if err != nil {
		response := APIResponse{Status: "error", Msg: "Gagal melakukan pengecekan pada nomor."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	if existingBobot != nil {
		response := APIResponse{Status: "error", Msg: "Nomor sudah ada."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
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
			response := APIResponse{Status: "error", Msg: "Gagal melakukan pengecekan pada parent nomor."}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}
		if parentBobot == nil {
			response := APIResponse{Status: "error", Msg: "Parent nomor tidak ditemukan."}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		bobot.ParentID = &parentBobot.ID
	}

	// Simpan bobot ke database
	if err := h.repo.CreateBobot(&bobot); err != nil {
		response := APIResponse{Status: "error", Msg: "Gagal menyimpan bobot."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Buat response
	bobotResponse := models.BobotResponse{
		Nama:  bobot.Nama,
		Nomor: bobot.Nomor,
	}

	response := APIResponse{Status: "success", Msg: "Bobot berhasil dibuat.", Data: bobotResponse}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *BobotHandler) GetAllBobots(w http.ResponseWriter, r *http.Request) {
	bobots, err := h.repo.GetAllBobots()
	if err != nil {
		response := APIResponse{Status: "error", Msg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
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

	response := APIResponse{Status: "success", Msg: "Data bobot berhasil diambil.", Data: bobotResponses}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
