package repository

import (
	"context"
	"dovenet/user-service/internal/domain"
	sql_repository "dovenet/user-service/internal/infrastructure/persistent/repository/sqlc"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	queries *sql_repository.Queries
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		queries: sql_repository.New(db.DB),
	}
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (*domain.User, error) {
	user, err := r.queries.UserGetByID(ctx, id)
	// TODO handle error
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Entity: domain.Entity{
			Id:        user.ID,
			CreatedAt: user.CreatedAt.Time,
			UpdatedAt: user.UpdatedAt.Time,
		},
		Email:       user.Username,
		Username:    user.Username,
		Firstname:   &user.Firstname.String,
		Lastname:    &user.Lastname.String,
		IsSuperuser: user.IsSuperuser.Bool,
		IsVerified:  user.IsVerified.Bool,
	}, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.queries.UserGetByEmail(ctx, email)
	// TODO handle error
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Entity: domain.Entity{
			Id:        user.ID,
			CreatedAt: user.CreatedAt.Time,
			UpdatedAt: user.UpdatedAt.Time,
		},
		Email:       user.Email,
		Username:    user.Username,
		Firstname:   &user.Firstname.String,
		Lastname:    &user.Lastname.String,
		IsSuperuser: user.IsSuperuser.Bool,
		IsVerified:  user.IsVerified.Bool,
	}, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := r.queries.UserGetByUsername(ctx, username)
	// TODO handle error
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Entity: domain.Entity{
			Id:        user.ID,
			CreatedAt: user.CreatedAt.Time,
			UpdatedAt: user.UpdatedAt.Time,
		},
		Email:       user.Username,
		Username:    user.Username,
		Firstname:   &user.Firstname.String,
		Lastname:    &user.Lastname.String,
		IsSuperuser: user.IsSuperuser.Bool,
		IsVerified:  user.IsVerified.Bool,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	result, err := r.queries.UserCreate(ctx, sql_repository.UserCreateParams{
		Email:       user.Email,
		Username:    user.Username,
		Firstname:   ToNullString(user.Firstname),
		Lastname:    ToNullString(user.Lastname),
		IsSuperuser: ToNullBool(user.IsSuperuser),
		IsVerified:  ToNullBool(user.IsVerified),
	})
	// TODO handle error
	if err != nil {
		return err
	}
	user.Entity.Id = result.ID
	user.Entity.CreatedAt = result.CreatedAt.Time
	user.Entity.UpdatedAt = result.UpdatedAt.Time
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	result, err := r.queries.UserUpdate(ctx, sql_repository.UserUpdateParams{
		ID:          user.Id,
		Email:       user.Email,
		Username:    user.Username,
		Firstname:   ToNullString(user.Firstname),
		Lastname:    ToNullString(user.Lastname),
		IsSuperuser: ToNullBool(user.IsSuperuser),
		IsVerified:  ToNullBool(user.IsVerified),
	})
	// TODO handle error
	if err != nil {
		return err
	}
	user.Entity.UpdatedAt = result.Time
	return nil
}

func (r *UserRepository) DeleteByID(ctx context.Context, id int32) error {
	_, err := r.queries.UserDeleteByID(ctx, id)
	// TODO handle error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteByEmail(ctx context.Context, email string) error {
	_, err := r.queries.UserDeleteByEmail(ctx, email)
	// TODO handle error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteByUsername(ctx context.Context, username string) error {
	_, err := r.queries.UserDeleteByUsername(ctx, username)
	// TODO handle error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	result, err := r.queries.UserExistsByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return result, nil

}

func (r *UserRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	result, err := r.queries.UserExistsByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	return result, nil

}
