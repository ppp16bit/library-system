package services

import (
	"fmt"
	"lib_backend/internal/model"
	"lib_backend/internal/repository"
	"log"

	"github.com/google/uuid"
)

type LoanService interface {
	CreateLoan(loan *model.Loan) (*model.Loan, error)
	GetLoanByID(id uuid.UUID) (*model.Loan, error)
	GetLoansByUserID(userID uuid.UUID) ([]model.Loan, error)
	GetLoansByBookID(bookID uuid.UUID) ([]model.Loan, error)
	ReturnBook(loanID uuid.UUID) (*model.Loan, error)
	DeleteLoan(id uuid.UUID) error
	GetAllLoans() ([]model.Loan, error)
}

type loanServiceImpl struct {
	loanRepo repository.LoanRepository
	userRepo repository.UserRepository
	bookRepo repository.BookRepository
}

func NewLoanService(loanRepo repository.LoanRepository, userRepo repository.UserRepository, bookRepo repository.BookRepository) LoanService {
	return &loanServiceImpl{loanRepo: loanRepo, userRepo: userRepo, bookRepo: bookRepo}
}

func (s *loanServiceImpl) CreateLoan(loan *model.Loan) (*model.Loan, error) {
	user, err := s.userRepo.GetUserByID(loan.UserID)

	if err != nil {
		return nil, fmt.Errorf("failed to check user existence for loan: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user with ID %s not found for loan", loan.UserID.String())
	}

	book, err := s.bookRepo.GetBookByID(loan.BookID)

	if err != nil {
		return nil, fmt.Errorf("failed to check book existence for loan: %w", err)
	}

	if book == nil {
		return nil, fmt.Errorf("book with ID %s not found for loan", loan.BookID.String())
	}

	if !book.Available {
		return nil, fmt.Errorf("book with ID %s is not available for loan", loan.BookID.String())
	}

	book.Available = false
	err = s.bookRepo.UpdateBook(book)

	if err != nil {
		return nil, fmt.Errorf("failed to update book availability after loan creation: %w", err)
	}

	err = s.loanRepo.CreateLoan(loan)

	if err != nil {
		return nil, fmt.Errorf("failed to create loan: %w", err)
	}

	return loan, nil
}

func (s *loanServiceImpl) GetLoanByID(id uuid.UUID) (*model.Loan, error) {
	loan, err := s.loanRepo.GetLoanByID(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get loan by ID: %w", err)
	}

	if loan == nil {
		return nil, fmt.Errorf("loan with ID %s not found", id.String())
	}

	return loan, nil
}

func (s *loanServiceImpl) GetLoansByUserID(userID uuid.UUID) ([]model.Loan, error) {
	loans, err := s.loanRepo.GetLoansByUserID(userID)

	if err != nil {
		return nil, fmt.Errorf("failed to get loans by user ID: %w", err)
	}

	return loans, nil
}

func (s *loanServiceImpl) GetLoansByBookID(bookID uuid.UUID) ([]model.Loan, error) {
	loans, err := s.loanRepo.GetLoansByBookID(bookID)

	if err != nil {
		return nil, fmt.Errorf("failed to get loans by book ID: %w", err)
	}

	return loans, nil
}

func (s *loanServiceImpl) ReturnBook(loanID uuid.UUID) (*model.Loan, error) {
	loan, err := s.loanRepo.GetLoanByID(loanID)

	if err != nil {
		return nil, fmt.Errorf("failed to get loan for return: %w", err)
	}

	if loan == nil {
		return nil, fmt.Errorf("loan with ID %s not found for return", loanID.String())
	}

	if loan.Returned {
		return nil, fmt.Errorf("loan with ID %s has already been returned", loanID.String())
	}

	loan.Returned = true
	err = s.loanRepo.UpdateLoan(loan)

	if err != nil {
		return nil, fmt.Errorf("failed to update loan status to returned: %w", err)
	}

	book, err := s.bookRepo.GetBookByID(loan.BookID)

	if err != nil {
		log.Printf("WARNING: failed to get book %s to update availability after return loan %s: %v", loan.BookID.String(), loanID.String(), err)
		return loan, nil
	}

	if book == nil {
		log.Printf("WARNING: book %s associated with returned loan %s not found cannot update availability", loan.BookID.String(), loanID.String())
		return loan, nil
	}
	book.Available = true
	err = s.bookRepo.UpdateBook(book)

	if err != nil {
		log.Printf("WARNING: failed to update book %s availability to true after return loan %s: %v", loan.BookID.String(), loanID.String(), err)
		return loan, nil
	}

	return loan, nil
}

func (s *loanServiceImpl) DeleteLoan(id uuid.UUID) error {
	err := s.loanRepo.DeleteLoan(id)

	if err != nil {
		return fmt.Errorf("failed to delete loan: %w", err)
	}

	return nil
}

func (s *loanServiceImpl) GetAllLoans() ([]model.Loan, error) {
	loans, err := s.loanRepo.GetAllLoans()

	if err != nil {
		return nil, fmt.Errorf("failed to get all loans: %w", err)
	}

	return loans, nil
}
