package domain

type User struct {
	Entity
	Email       string  `json:"email"`
	Username    string  `json:"username"`
	Firstname   *string `json:"firstname"`
	Lastname    *string `json:"lastname"`
	IsSuperuser bool    `json:"is_superuser"`
	IsVerified  bool    `json:"is_verified"`
}
