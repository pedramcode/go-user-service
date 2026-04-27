package domain

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
