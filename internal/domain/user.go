package domain

import "context"

type User struct {
	Entity
	Email       string  `json:"email"`
	Username    string  `json:"username"`
	Firstname   *string `json:"firstname"`
	Lastname    *string `json:"lastname"`
	IsSuperuser bool    `json:"is_superuser"`
	IsVerified  bool    `json:"is_verified"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id int32) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	DeleteByID(ctx context.Context, id int32) error
	DeleteByEmail(ctx context.Context, email string) error
	DeleteByUsername(ctx context.Context, username string) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}
