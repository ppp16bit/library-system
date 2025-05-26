package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"lib_backend/internal/model"

	"github.com/google/uuid"
)

type LoanRepository interface {
	CreateLoan(loan *model.Loan) error
	GetLoanByID(id uuid.UUID) (*model.Loan, error)
	GetLoansByUserID(userID uuid.UUID) ([]model.Loan, error)
	GetLoansByBookID(bookID uuid.UUID) ([]model.Loan, error)
	UpdateLoan(loan *model.Loan) error
	DeleteLoan(id uuid.UUID) error
	GetAllLoans() ([]model.Loan, error)
}

type loanRepositoryImpl struct {
	db *sql.DB
}

func NewLoanRepository(db *sql.DB) LoanRepository {
	return &loanRepositoryImpl{db: db}
}

func (r *loanRepositoryImpl) CreateLoan(loan *model.Loan) error {
	loan.ID = uuid.New()
	loan.LoanedAt = time.Now()

	query := `INSERT INTO loans (id, user_id, book_id, loaned_at, returned) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, loan.ID, loan.UserID, loan.BookID, loan.LoanedAt, loan.Returned)

	if err != nil {
		return fmt.Errorf("failed to create loan for user ID %s and book ID %s: %w", loan.UserID.String(), loan.BookID.String(), err)
	}

	return nil
}

func (r *loanRepositoryImpl) GetLoanByID(id uuid.UUID) (*model.Loan, error) {
	loan := &model.Loan{}
	query := `SELECT id, user_id, book_id, loaned_at, returned FROM loans WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&loan.ID, &loan.UserID, &loan.BookID, &loan.LoanedAt, &loan.Returned)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get loan by ID %s: %w", id.String(), err)
	}

	return loan, nil
}

func (r *loanRepositoryImpl) GetLoansByUserID(userID uuid.UUID) ([]model.Loan, error) {
	query := `SELECT id, user_id, book_id, loaned_at, returned FROM loans WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to get loans for user ID %s: %w", userID.String(), err)
	}
	defer func() {

		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("ERROR: failed to close rows after getting loans by user ID %s: %v", userID.String(), closeErr)
		}
	}()

	var loans []model.Loan

	for rows.Next() {
		loan := model.Loan{}

		if err := rows.Scan(&loan.ID, &loan.UserID, &loan.BookID, &loan.LoanedAt, &loan.Returned); err != nil {
			return nil, fmt.Errorf("failed to scan loan row for user ID %s: %w", userID.String(), err)
		}
		loans = append(loans, loan)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during loan rows iteration for user ID %s: %w", userID.String(), err)
	}

	return loans, nil
}

func (r *loanRepositoryImpl) GetLoansByBookID(bookID uuid.UUID) ([]model.Loan, error) {
	query := `SELECT id, user_id, book_id, loaned_at, returned FROM loans WHERE book_id = $1`
	rows, err := r.db.Query(query, bookID)

	if err != nil {
		return nil, fmt.Errorf("failed to get loans for book ID %s: %w", bookID.String(), err)
	}
	defer func() {

		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("ERROR: failed to close rows after getting loans by book ID %s: %v", bookID.String(), closeErr)
		}
	}()

	var loans []model.Loan

	for rows.Next() {
		loan := model.Loan{}

		if err := rows.Scan(&loan.ID, &loan.UserID, &loan.BookID, &loan.LoanedAt, &loan.Returned); err != nil {
			return nil, fmt.Errorf("failed to scan loan row for book ID %s: %w", bookID.String(), err)
		}

		loans = append(loans, loan)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during loan rows iteration for book ID %s: %w", bookID.String(), err)
	}

	return loans, nil
}

func (r *loanRepositoryImpl) UpdateLoan(loan *model.Loan) error {
	query := `UPDATE loans SET user_id = $2, book_id = $3, loaned_at = $4, returned = $5 WHERE id = $1`
	res, err := r.db.Exec(query, loan.ID, loan.UserID, loan.BookID, loan.LoanedAt, loan.Returned)

	if err != nil {
		return fmt.Errorf("failed to execute update query for loan ID %s: %w", loan.ID.String(), err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return fmt.Errorf("failed to check rows affected after updating loan ID %s: %w", loan.ID.String(), err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("loan with ID %s not found for update or no changes were made", loan.ID)
	}

	return nil
}

func (r *loanRepositoryImpl) DeleteLoan(id uuid.UUID) error {
	query := `DELETE FROM loans WHERE id = $1`
	res, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed to execute delete query for loan ID %s: %w", id.String(), err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return fmt.Errorf("failed to check rows affected after deleting loan ID %s: %w", id.String(), err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("loan with ID %s not found for deletion", id)
	}

	return nil
}

func (r *loanRepositoryImpl) GetAllLoans() ([]model.Loan, error) {
	query := `SELECT id, user_id, book_id, loaned_at, returned FROM loans`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("failed to query all loans from database: %w", err)
	}
	defer func() {

		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("ERROR: failed to close rows after getting all loans: %v", closeErr)
		}
	}()

	var loans []model.Loan

	for rows.Next() {
		loan := model.Loan{}

		if err := rows.Scan(&loan.ID, &loan.UserID, &loan.BookID, &loan.LoanedAt, &loan.Returned); err != nil {
			return nil, fmt.Errorf("failed to scan loan row into struct: %w", err)
		}
		loans = append(loans, loan)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration over loan rows: %w", err)
	}

	return loans, nil
}
