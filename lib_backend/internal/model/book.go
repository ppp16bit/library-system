package model

import "github.com/google/uuid"

type Book struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Isbn      string    `json:"isbn"`
	Available bool      `json:"available"`
}
