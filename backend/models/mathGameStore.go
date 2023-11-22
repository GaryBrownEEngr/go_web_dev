package models

import (
	"time"
)

type MathGameData struct {
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
	JsonData  string    `json:"json_data"`
}

type MathGameStore interface {
	Read(username string) (*MathGameData, error)
	Write(in *MathGameData) error
}
