package port

import (
	"context"
	"dovenet/user-service/internal/domain"
)

type CredentialRepository interface {
	GetByID(ctx context.Context, id int32) (*domain.Credential, error)
	GetByUserTypeKey(ctx context.Context, userID int32, ctype domain.CredentialType, key string) (*domain.Credential, error)
	Create(ctx context.Context, credential *domain.Credential) error
	Update(ctx context.Context, credential *domain.Credential) error
	DeleteByID(ctx context.Context, id int32) error
}
