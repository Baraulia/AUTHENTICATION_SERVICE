package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email" `
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUser struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"password"`
}

type UpdateUser struct {
	Email       string `json:"email" validate:"email"`
	OldPassword string `json:"old_password" validate:"password"`
	NewPassword string `json:"new_password" validate:"password"`
}

// Users array of User type

type Users []User

var PasswordComposition = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	"0123456789")
