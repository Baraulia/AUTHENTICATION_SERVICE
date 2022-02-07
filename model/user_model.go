package model

import "time"

type User struct {
	ID             int64     `json:"id"       sql:"id"`
	Email          string    `json:"email" validate:"required" sql:"email"`
	Password       string    `json:"password" validate:"required" sql:"password"`
	Activated      string    `json:"activated" example:"false" sql:"activated"`
	CreatedAt      time.Time `json:"createdAt" sql:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" sql:"updated_at"`
}

// Users array of User type

type Users []User
