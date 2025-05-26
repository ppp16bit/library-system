package model

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Registration string    `json:"registration"`
	Email        string    `json:"email"`
}
