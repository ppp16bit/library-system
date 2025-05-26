package services

import (
	"fmt"
	"lib_backend/internal/model"
	"lib_backend/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByID(id uuid.UUID) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id uuid.UUID) error
	GetAllUsers() ([]model.User, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	existingUser, err := s.userRepo.GetUserByEmail(user.Email)

	if err != nil {
		return nil, fmt.Errorf("failed to check for existing user by email: %w", err)
	}

	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	err = s.userRepo.CreateUser(user)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *userServiceImpl) GetUserByID(id uuid.UUID) (*model.User, error) {
	user, err := s.userRepo.GetUserByID(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user with ID %s not found", id.String())
	}

	return user, nil
}

func (s *userServiceImpl) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user with email %s not found", email)
	}

	return user, nil
}

func (s *userServiceImpl) UpdateUser(user *model.User) (*model.User, error) {
	existingUser, err := s.userRepo.GetUserByID(user.ID) // verifica se o user existe antes do update

	if err != nil {
		return nil, fmt.Errorf("failed to check for existing user before update: %w", err)
	}

	if existingUser == nil {
		return nil, fmt.Errorf("user with ID %s not found for update", user.ID.String())
	}

	err = s.userRepo.UpdateUser(user)

	if err != nil {
		return nil, fmt.Errorf("failed to update user %w", err)
	}

	return user, nil
}

func (s *userServiceImpl) DeleteUser(id uuid.UUID) error {
	err := s.userRepo.DeleteUser(id)

	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (s *userServiceImpl) GetAllUsers() ([]model.User, error) {
	users, err := s.userRepo.GetAllUsers()

	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}
