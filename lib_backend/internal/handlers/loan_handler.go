package handler

import (
	"log"
	"net/http"

	"lib_backend/internal/dto"
	"lib_backend/internal/model"
	"lib_backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoanHandler struct {
	loanService services.LoanService
}

func NewLoanHandler(s services.LoanService) *LoanHandler {
	return &LoanHandler{loanService: s}
}

func (h *LoanHandler) CreateLoan(c *gin.Context) {
	var request dto.LoanRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	parsedUserID := uuid.MustParse(request.UserID)
	parsedBookID := uuid.MustParse(request.BookID)

	loanToCreate := &model.Loan{
		UserID:   parsedUserID,
		BookID:   parsedBookID,
		Returned: false,
		LoanedAt: model.DefaultLoanedAt(),
	}

	createdLoan, err := h.loanService.CreateLoan(loanToCreate)

	if err != nil {
		switch err.Error() {
		case "user with ID " + parsedUserID.String() + " not found for loan":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		case "book with ID " + parsedBookID.String() + " is not available for loan":
			c.JSON(http.StatusConflict, gin.H{"error": "Book is not available for loan"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan", "details": err.Error()})
		}
		return
	}
	c.JSON(http.StatusCreated, createdLoan)
}

func (h *LoanHandler) GetLoanByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID format", "details": err.Error()})
		return
	}

	loan, err := h.loanService.GetLoanByID(id)

	if err != nil {
		log.Printf("ERROR: GetLoanByID service failed for ID %s: %v", id.String(), err)

		if err.Error() == "service: loan with ID "+id.String()+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loan", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (h *LoanHandler) GetLoansByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format", "details": err.Error()})
		return
	}

	loans, err := h.loanService.GetLoansByUserID(userID)

	if err != nil {
		log.Printf("ERROR: GetLoansByUserID service failed for user ID %s: %v", userID.String(), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loans for user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}

func (h *LoanHandler) GetLoansByBookID(c *gin.Context) {
	bookIDStr := c.Param("book_id")
	bookID, err := uuid.Parse(bookIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format", "details": err.Error()})
		return
	}

	loans, err := h.loanService.GetLoansByBookID(bookID)

	if err != nil {
		log.Printf("ERROR: GetLoansByBookID service failed for book ID %s: %v", bookID.String(), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loans for book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}

func (h *LoanHandler) ReturnBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID format", "details": err.Error()})
		return
	}

	returnedLoan, err := h.loanService.ReturnBook(id)

	if err != nil {
		log.Printf("ERROR: ReturnBook service failed for ID %s: %v", id.String(), err)

		switch err.Error() {
		case "service: loan with ID " + id.String() + " not found for return":
			c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		case "service: loan with ID " + id.String() + " has already been returned":
			c.JSON(http.StatusConflict, gin.H{"error": "Loan already returned"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to return book", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, returnedLoan)
}

func (h *LoanHandler) DeleteLoan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID format", "details": err.Error()})
		return
	}

	err = h.loanService.DeleteLoan(id)

	if err != nil {
		log.Printf("ERROR: DeleteLoan service failed for ID %s: %v", id.String(), err)

		if err.Error() == "service: loan with ID "+id.String()+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete loan", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *LoanHandler) GetAllLoans(c *gin.Context) {
	loans, err := h.loanService.GetAllLoans()

	if err != nil {
		log.Printf("ERROR: GetAllLoans service failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loans", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}
