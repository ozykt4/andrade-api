package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type UserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}

func (u *UserReq) ToUser() *User {
	return &User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (ur *User) ToUserRes() *UserRes {
	return &UserRes{
		ID:        ur.ID,
		Name:      ur.Name,
		Email:     ur.Email,
		CreatedAT: ur.CreatedAt,
		UpdatedAT: ur.UpdatedAt,
	}
}
