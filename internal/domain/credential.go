package domain

import "context"

type CredentialType string

const (
	Password CredentialType = "password"
	Google   CredentialType = "google"
	LinkedIn CredentialType = "linkedin"
	Facebook CredentialType = "facebook"
)

type Credential struct {
	Entity
	UserID int32
	Type   CredentialType
	Key    string
	Value  string
}

type CredentialRepository interface {
	GetByID(ctx context.Context, id int32) (*Credential, error)
	GetByUserTypeKey(ctx context.Context, userID int32, ctype CredentialType, key string) (*Credential, error)
	Create(ctx context.Context, credential *Credential) error
	Update(ctx context.Context, credential *Credential) error
	DeleteByID(ctx context.Context, id int32) error
}
