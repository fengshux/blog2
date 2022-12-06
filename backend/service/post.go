package service

import (
	"context"
	"errors"

	"github.com/fengshux/blog2/backend/model"
	"gorm.io/gorm"
)

type Post struct {
	BaseService
}

func NewPost(baseService BaseService) *Post {
	return &Post{
		baseService,
	}
}

func (p *Post) List(ctx context.Context, opts *model.SQLOption) ([]model.Post, error) {

	var posts []model.Post

	query := p.DB(ctx).Table("post")
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

	result := query.Find(&posts)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return posts, nil
}

func (p *Post) Count(ctx context.Context) (int64, error) {

	var count int64 = 0
	result := p.DB(ctx).Table("post").Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (p *Post) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	result := p.DB(ctx).Table("post").Create(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (p *Post) FindOne(ctx context.Context, query *model.Post) (*model.Post, error) {
	post := &model.Post{}
	result := p.DB(ctx).Table("post").Last(post, query.ID)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return post, nil
}
