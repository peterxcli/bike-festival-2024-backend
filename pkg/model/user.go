package model

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   string `json:"id" gorm:"type:varchar(36);primary_key"`
	Name string `json:"name" gorm:"type:varchar(255);index"`
}

func (u *User) BeforeCreate(*gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
	CreateAccessToken(ctx context.Context, user *User, secret string, expiry int64) (accessToken string, err error)
	CreateRefreshToken(ctx context.Context, user *User, secret string, expiry int64) (refreshToken string, err error)
	VerifyRefreshToken(ctx context.Context, refreshToken string, secret string) (user *User, err error)
	Logout(ctx context.Context, token *string, secrect string) error
}
