package dto

type LoanRequest struct {
	UserID string `json:"userId" binding:"required,uuid"`
	BookID string `json:"bookId" binding:"required,uuid"`
}
