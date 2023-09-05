package models

import (
	"time"
)

type Flag struct {
	ID        int64     `json:"id"`
	Flag      string    `json:"flag"`
	Value     float64   `json:"value"`
	Mounth    string    `json:"mounth"`
	MountNum  int64     `json:"mount_num"`
	Year      int64     `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
