package service

import (
	"context"
	"errors"

	"github.com/fengshux/blog2/backend/model"
	"gorm.io/gorm"
)

type User struct {
	BaseService
}

func NewUser(baseService BaseService) *User {
	return &User{
		baseService,
	}
}

func (u *User) List(ctx context.Context, opts *model.SQLOption) ([]model.User, error) {

	var users []model.User

	query := u.DB(ctx).Table("user")
	if opts != nil {
		if opts.Limit != 0 {
			query.Limit(opts.Limit)
		}

		if opts.Offset != 0 {
			query.Offset(opts.Offset)
		}

		if opts.OrderBy != "" {
			query.Order(opts.OrderBy)
		}
	}

	result := query.Find(&users)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return users, nil
}

func (u *User) Count(ctx context.Context) (int64, error) {

	var count int64 = 0
	result := u.DB(ctx).Table("user").Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (u *User) Create(ctx context.Context, user *model.User) (*model.User, error) {
	result := u.DB(ctx).Table("user").Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
