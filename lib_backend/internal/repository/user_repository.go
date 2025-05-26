package repository

import (
	"database/sql"
	"fmt"
	"log"

	"lib_backend/internal/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id uuid.UUID) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uuid.UUID) error
	GetAllUsers() ([]model.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) CreateUser(user *model.User) error {
	user.ID = uuid.New()

	query := `INSERT INTO users (id, name, registration, email) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.ID, user.Name, user.Registration, user.Email)

	if err != nil {
		return fmt.Errorf("failed to create user with email %s and registration %s: %w", user.Email, user.Registration, err)
	}

	return nil
}

func (r *userRepositoryImpl) GetUserByID(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, name, registration, email FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Registration, &user.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by ID %s: %w", id.String(), err)
	}

	return user, nil
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, name, registration, email FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Registration, &user.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by email %s: %w", email, err)
	}

	return user, nil
}

func (r *userRepositoryImpl) UpdateUser(user *model.User) error {
	query := `UPDATE users SET name = $2, registration = $3, email = $4 WHERE id = $1`
	res, err := r.db.Exec(query, user.ID, user.Name, user.Registration, user.Email)

	if err != nil {
		return fmt.Errorf("failed to execute update query for user ID %s: %w", user.ID.String(), err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return fmt.Errorf("failed to check rows affected after updating user ID %s: %w", user.ID.String(), err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found for update or no changes were made", user.ID)
	}

	return nil
}

func (r *userRepositoryImpl) DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	res, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed to execute delete query for user ID %s: %w", id.String(), err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return fmt.Errorf("failed to check rows affected after deleting user ID %s: %w", id.String(), err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found for deletion", id)
	}

	return nil
}

func (r *userRepositoryImpl) GetAllUsers() ([]model.User, error) {
	query := `SELECT id, name, registration, email FROM users`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("failed to query all users from database: %w", err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("ERROR: failed to close rows after getting all users: %v", closeErr)
		}
	}()

	var users []model.User

	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Registration, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user row into struct: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration over user rows: %w", err)
	}

	return users, nil
}
