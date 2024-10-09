package models

type Bobot struct {
	ID       int    `db:"id" json:"id"`
	ParentID *int   `db:"parent_id" json:"parent_id"`
	Nama     string `db:"nama" json:"nama"`
	Nomor    string `db:"nomor" json:"nomor"`
}

type BobotSpec struct {
	Nama  string `json:"nama"`
	Nomor string `json:"nomor"`
}

type BobotResponse struct {
	Nama  string `json:"nama"`
	Nomor string `json:"nomor"`
}
