package port

import (
	"context"
	"dovenet/user-service/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int32) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	DeleteByID(ctx context.Context, id int32) error
	DeleteByEmail(ctx context.Context, email string) error
	DeleteByUsername(ctx context.Context, username string) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}
