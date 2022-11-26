package service

import (
	"context"
	"errors"

	"github.com/fengshux/blog2/backend/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) List(ctx context.Context) ([]model.User, error) {

	var users []model.User

	result := u.db.Table("users").Find(&users)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return users, nil
}
