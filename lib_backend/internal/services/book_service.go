package services

import (
	"fmt"
	"lib_backend/internal/model"
	"lib_backend/internal/repository"

	"github.com/google/uuid"
)

type BookService interface {
	CreateBook(book *model.Book) (*model.Book, error)
	GetBookByID(id uuid.UUID) (*model.Book, error)
	GetBookByISBN(isbn string) (*model.Book, error)
	UpdateBook(book *model.Book) (*model.Book, error)
	DeleteBook(id uuid.UUID) error
	GetAllBooks() ([]model.Book, error)
}

type bookServiceImpl struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookServiceImpl{bookRepo: bookRepo}
}

func (s *bookServiceImpl) CreateBook(book *model.Book) (*model.Book, error) {
	existingBook, err := s.bookRepo.GetBookByISBN(book.Isbn)

	if err != nil {
		return nil, fmt.Errorf("failed to check for existing book by ISBN: %w", err)
	}

	if existingBook != nil {
		return nil, fmt.Errorf("book with ISBN %s already exists", book.Isbn)
	}

	err = s.bookRepo.CreateBook(book)

	if err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	return book, nil
}

func (s *bookServiceImpl) GetBookByID(id uuid.UUID) (*model.Book, error) {
	book, err := s.bookRepo.GetBookByID(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get book by ID: %w", err)
	}

	if book == nil {
		return nil, fmt.Errorf("book with ID %s not found", id.String())
	}

	return book, nil
}

func (s *bookServiceImpl) GetBookByISBN(isbn string) (*model.Book, error) {
	book, err := s.bookRepo.GetBookByISBN(isbn)

	if err != nil {
		return nil, fmt.Errorf("failed to get book by ISBN: %w", err)
	}

	if book == nil {
		return nil, fmt.Errorf("book with ISBN %s not found", isbn)
	}

	return book, nil
}

func (s *bookServiceImpl) UpdateBook(book *model.Book) (*model.Book, error) {
	existingBook, err := s.bookRepo.GetBookByID(book.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to check for existing book before update: %w", err)
	}

	if existingBook == nil {
		return nil, fmt.Errorf("book with ID %s not found for update", book.ID.String())
	}

	err = s.bookRepo.UpdateBook(book)

	if err != nil {
		return nil, fmt.Errorf("failed to update book: %w", err)
	}

	return book, nil
}

func (s *bookServiceImpl) DeleteBook(id uuid.UUID) error {
	err := s.bookRepo.DeleteBook(id)

	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	return nil
}

func (s *bookServiceImpl) GetAllBooks() ([]model.Book, error) {
	books, err := s.bookRepo.GetAllBooks()
	if err != nil {
		return nil, fmt.Errorf("failed to get all books: %w", err)
	}
	return books, nil
}
