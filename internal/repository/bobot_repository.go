package repository

import (
	"go-tc-plnsmrng/internal/models"

	"github.com/jmoiron/sqlx"
)

type BobotRepository struct {
	db *sqlx.DB
}

func NewBobotRepository(db *sqlx.DB) *BobotRepository {
	return &BobotRepository{db: db}
}

func (r *BobotRepository) CreateBobot(bobot *models.Bobot) error {
	query := `INSERT INTO bobot (parent_id, nama, nomor) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, bobot.ParentID, bobot.Nama, bobot.Nomor).Scan(&bobot.ID)
}

func (r *BobotRepository) GetAllBobots() ([]models.Bobot, error) {
	var bobots []models.Bobot
	query := "SELECT id, parent_id, nama, nomor FROM bobot ORDER BY nomor ASC"
	err := r.db.Select(&bobots, query)
	return bobots, err
}
