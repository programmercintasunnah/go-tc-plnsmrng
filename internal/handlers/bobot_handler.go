package handlers

import (
	"encoding/json"
	"go-tc-plnsmrng/internal/models"
	"go-tc-plnsmrng/internal/repository"
	"net/http"
)

type BobotHandler struct {
	repo *repository.BobotRepository
}

func NewBobotHandler(repo *repository.BobotRepository) *BobotHandler {
	return &BobotHandler{repo: repo}
}

func (h *BobotHandler) CreateBobot(w http.ResponseWriter, r *http.Request) {
	var bobot models.Bobot
	if err := json.NewDecoder(r.Body).Decode(&bobot); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateBobot(&bobot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bobot)
}

func (h *BobotHandler) GetAllBobots(w http.ResponseWriter, r *http.Request) {
	bobots, err := h.repo.GetAllBobots()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bobots)
}
