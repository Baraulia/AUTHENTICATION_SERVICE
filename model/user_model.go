package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email" `
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUser struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	RoleId   int    `json:"role_id" binding:"required" validate:"roleId"`
	Password string `json:"password" validate:"password"`
}

type UpdateUser struct {
	Email       string `json:"email" validate:"email"`
	OldPassword string `json:"old_password" binding:"required" validate:"password"`
	NewPassword string `json:"new_password" binding:"required" validate:"password"`
}
type ResponseUser struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type MockUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email" `
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

// Users array of User type

type Users []User

var PasswordNumber = []rune("0123456789")

var PasswordUpper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var PasswordLower = []rune("abcdefghijklmnopqrstuvwxyz")

var PasswordSpecial = []rune("@#%&!$")

var PasswordComposition = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	"0123456789" +
	"@#%&!$")
