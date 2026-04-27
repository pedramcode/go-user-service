package application

import (
	"context"
	"dovenet/user-service/internal/domain"
)

type UserService struct {
	userRepo       domain.UserRepository
	credentialRepo domain.CredentialRepository
}

func NewUserService(
	userRepo domain.UserRepository,
	credentialRepo domain.CredentialRepository) *UserService {
	return &UserService{
		userRepo:       userRepo,
		credentialRepo: credentialRepo,
	}
}

func (s *UserService) CreateSuperuser(ctx context.Context, username string, email string, password string) (*domain.User, error) {
	user := domain.User{
		Email:       email,
		Username:    username,
		IsSuperuser: true,
		IsVerified:  true,
	}
	// TODO handle error
	if err := s.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	credUsername := domain.Credential{
		UserID: user.Entity.Id,
		Type:   domain.Password,
		Key:    "username",
		Value:  password,
	}
	// TODO handle error
	if err := s.credentialRepo.Create(ctx, &credUsername); err != nil {
		return nil, err
	}

	credEmail := domain.Credential{
		UserID: user.Entity.Id,
		Type:   domain.Password,
		Key:    "email",
		Value:  password,
	}
	// TODO handle error
	if err := s.credentialRepo.Create(ctx, &credEmail); err != nil {
		return nil, err
	}

	return &user, nil
}
