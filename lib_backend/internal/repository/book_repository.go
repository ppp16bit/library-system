package repository

import (
	"database/sql"
	"fmt"
	"log"

	"lib_backend/internal/model"

	"github.com/google/uuid"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetBookByID(id uuid.UUID) (*model.Book, error)
	GetBookByISBN(isbn string) (*model.Book, error)
	UpdateBook(book *model.Book) error
	DeleteBook(id uuid.UUID) error
	GetAllBooks() ([]model.Book, error)
}

type bookRepositoryImpl struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepositoryImpl{db: db}
}

func (r *bookRepositoryImpl) CreateBook(book *model.Book) error {
	book.ID = uuid.New()

	query := `INSERT INTO books (id, title, author, isbn, available) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, book.ID, book.Title, book.Author, book.Isbn, book.Available)

	if err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}

	return nil
}

func (r *bookRepositoryImpl) GetBookByID(id uuid.UUID) (*model.Book, error) {
	book := &model.Book{}
	query := `SELECT id, title, author, isbn, available FROM books WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Isbn, &book.Available)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get book by ID %s: %w", id.String(), err)
	}

	return book, nil
}

func (r *bookRepositoryImpl) GetBookByISBN(isbn string) (*model.Book, error) {
	book := &model.Book{}
	query := `SELECT id, title, author, isbn, available FROM books WHERE isbn = $1`
	err := r.db.QueryRow(query, isbn).Scan(&book.ID, &book.Title, &book.Author, &book.Isbn, &book.Available)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get book by ISBN %s: %w", isbn, err)
	}

	return book, nil
}

func (r *bookRepositoryImpl) UpdateBook(book *model.Book) error {
	query := `UPDATE books SET title = $2, author = $3, isbn = $4, available = $5 WHERE id = $1`
	res, err := r.db.Exec(query, book.ID, book.Title, book.Author, book.Isbn, book.Available)

	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return fmt.Errorf("failed to check rows affected after update: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("book with ID %s not found for update", book.ID)
	}

	return nil
}

func (r *bookRepositoryImpl) DeleteBook(id uuid.UUID) error {
	query := `DELETE FROM books WHERE id = $1`
	res, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return fmt.Errorf("failed to check rows affected after deletion: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("book with ID %s not found for deletion", id)
	}

	return nil
}

func (r *bookRepositoryImpl) GetAllBooks() ([]model.Book, error) {
	query := `SELECT id, title, author, isbn, available FROM books`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("failed to get all books: %w", err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("error closing rows in GetAllBooks: %v", closeErr)
		}
	}()

	var books []model.Book

	for rows.Next() {
		book := model.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Isbn, &book.Available); err != nil {
			return nil, fmt.Errorf("failed to scan book row: %w", err)
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during book rows iteration: %w", err)
	}

	return books, nil
}
