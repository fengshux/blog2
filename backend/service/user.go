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

func (u *User) List(ctx context.Context, where model.SQLWhere, opts *model.SQLOption) ([]model.User, error) {

	var users []model.User

	query := u.DB(ctx).Table("user")

	if len(where) != 0 {
		statement, params := where.ToGormHere()
		query = query.Where(statement, params...)
	}

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

func (u *User) Create(ctx context.Context, user *model.FullUser) (*model.FullUser, error) {
	result := u.DB(ctx).Table("user").Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *User) FindOne(ctx context.Context, query *model.User) (*model.User, error) {
	user := &model.User{}
	result := u.DB(ctx).Model(&model.User{}).Where(query).Last(user)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, nil
}

func (u *User) FindOneFullUser(ctx context.Context, query *model.User) (*model.FullUser, error) {
	fullUser := &model.FullUser{}
	result := u.DB(ctx).Model(&model.User{}).Where(query).Last(fullUser)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return fullUser, nil
}
