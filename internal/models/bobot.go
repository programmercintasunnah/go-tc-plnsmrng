package models

type Bobot struct {
	ID       int    `db:"id"`
	ParentID *int   `db:"parent_id"`
	Nama     string `db:"nama"`
	Nomor    string `db:"nomor"`
}
