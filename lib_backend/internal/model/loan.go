package model

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID         uuid.UUID  `json:"id"`
	UserID     uuid.UUID  `json:"user_id"`
	BookID     uuid.UUID  `json:"book_id"`
	LoanedAt   time.Time  `json:"loaned_at"`
	Returned   bool       `json:"returned"`
	ReturnedAt *time.Time `json:"returned_at,omitempty"`
}
