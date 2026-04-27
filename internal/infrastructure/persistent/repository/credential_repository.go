package repository

import (
	"context"
	"dovenet/user-service/internal/domain"
	sql_repository "dovenet/user-service/internal/infrastructure/persistent/repository/sqlc"

	"github.com/jmoiron/sqlx"
)

type CredentialRepository struct {
	queries *sql_repository.Queries
}

func NewCredentialRepository(db *sqlx.DB) *CredentialRepository {
	return &CredentialRepository{
		queries: sql_repository.New(db.DB),
	}
}

func (r *CredentialRepository) GetByID(ctx context.Context, id int32) (*domain.Credential, error) {
	result, err := r.queries.CredentialGetByID(ctx, id)
	// TODO handle error
	if err != nil {
		return nil, err
	}
	return &domain.Credential{
		Entity: domain.Entity{
			Id:        result.ID,
			CreatedAt: result.CreatedAt.Time,
			UpdatedAt: result.CreatedAt.Time,
			DeletedAt: result.DeletedAt.Time,
		},
		UserID: result.UserID,
		Type:   domain.CredentialType(result.Type),
		Key:    result.Key,
		Value:  result.Value,
	}, nil
}

func (r *CredentialRepository) GetByUserTypeKey(ctx context.Context, userID int32, ctype domain.CredentialType, key string) (*domain.Credential, error) {
	result, err := r.queries.CredentialGetByUserTypeKey(ctx, sql_repository.CredentialGetByUserTypeKeyParams{
		UserID: userID,
		Type:   string(ctype),
		Key:    key,
	})
	// TODO handle error
	if err != nil {
		return nil, err
	}
	return &domain.Credential{
		Entity: domain.Entity{
			Id:        result.ID,
			CreatedAt: result.CreatedAt.Time,
			UpdatedAt: result.CreatedAt.Time,
			DeletedAt: result.DeletedAt.Time,
		},
		UserID: result.UserID,
		Type:   domain.CredentialType(result.Type),
		Key:    result.Key,
		Value:  result.Value,
	}, nil
}

func (r *CredentialRepository) Create(ctx context.Context, credential *domain.Credential) error {
	result, err := r.queries.CredentialCreate(ctx, sql_repository.CredentialCreateParams{
		UserID: credential.UserID,
		Type:   string(credential.Type),
		Key:    credential.Key,
		Value:  credential.Value,
	})
	// TODO handle error
	if err != nil {
		return err
	}
	credential.Entity.Id = result.ID
	credential.Entity.CreatedAt = result.CreatedAt.Time
	credential.Entity.UpdatedAt = result.UpdatedAt.Time
	return nil
}

func (r *CredentialRepository) Update(ctx context.Context, credential *domain.Credential) error {
	result, err := r.queries.CredentialUpdate(ctx, sql_repository.CredentialUpdateParams{
		ID:     credential.Id,
		UserID: credential.UserID,
		Type:   string(credential.Type),
		Key:    credential.Key,
		Value:  credential.Value,
	})
	// TODO handle error
	if err != nil {
		return err
	}
	credential.Entity.UpdatedAt = result.Time
	return nil
}

func (r *CredentialRepository) DeleteByID(ctx context.Context, id int32) error {
	_, err := r.queries.CredentialDeleteByID(ctx, id)
	// TODO handle error
	if err != nil {
		return err
	}
	return nil
}
