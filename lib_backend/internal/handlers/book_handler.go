package handler

import (
	"log"
	"net/http"

	"lib_backend/internal/model"
	"lib_backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookHandler struct {
	bookService services.BookService
}

func NewBookHandler(s services.BookService) *BookHandler {
	return &BookHandler{bookService: s}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book model.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	createdBook, err := h.bookService.CreateBook(&book)

	if err != nil {
		log.Printf("ERROR: CreateBook service failed: %v", err)

		if err.Error() == "service: book with ISBN "+book.Isbn+" already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Book with this ISBN already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdBook)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format", "details": err.Error()})
		return
	}

	book, err := h.bookService.GetBookByID(id)

	if err != nil {
		log.Printf("ERROR: GetBookByID service failed for ID %s: %v", id.String(), err)

		if err.Error() == "service: book with ID "+id.String()+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) GetBookByISBN(c *gin.Context) {
	isbn := c.Query("isbn")

	if isbn == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ISBN parameter is required"})
		return
	}

	book, err := h.bookService.GetBookByISBN(isbn)

	if err != nil {
		log.Printf("ERROR: GetBookByISBN service failed for ISBN %s: %v", isbn, err)

		if err.Error() == "service: book with ISBN "+isbn+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format", "details": err.Error()})
		return
	}

	var book model.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	book.ID = id

	updatedBook, err := h.bookService.UpdateBook(&book)

	if err != nil {
		log.Printf("ERROR: UpdateBook service failed for ID %s: %v", id.String(), err)

		if err.Error() == "service: book with ID "+id.String()+" not found for update" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format", "details": err.Error()})
		return
	}

	err = h.bookService.DeleteBook(id)

	if err != nil {
		log.Printf("ERROR: DeleteBook service failed for ID %s: %v", id.String(), err)

		if err.Error() == "service: book with ID "+id.String()+" not found for update" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks()

	if err != nil {
		log.Printf("ERROR: GetAllBooks service failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
