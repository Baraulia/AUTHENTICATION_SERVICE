package model

type AuthUser struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"password"`
}
