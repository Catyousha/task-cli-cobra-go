package db

import "time"

var dbLoc string

type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status" validate:"oneof=todo in-progress done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Initialize(db string) (string, error) {
	dbLoc = db
	return dbLoc, nil
}