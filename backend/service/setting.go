package service

import (
	"context"
	"errors"

	"github.com/fengshux/blog2/backend/model"
	"gorm.io/gorm"
)

type Setting struct {
	BaseService
}

func NewSetting(baseService BaseService) *Setting {
	return &Setting{
		baseService,
	}
}

func (s *Setting) List(ctx context.Context, opts *model.SQLOption) ([]model.Setting, error) {

	var settings []model.Setting

	query := s.DB(ctx).Table("setting")

	result := query.Find(&settings)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return settings, nil
}

func (s *Setting) Create(ctx context.Context, setting *model.Setting) (*model.Setting, error) {
	result := s.DB(ctx).Table("setting").Create(setting)
	if result.Error != nil {
		return nil, result.Error
	}
	return setting, nil
}

func (s *Setting) FindOne(ctx context.Context, key string) (*model.Setting, error) {
	setting := &model.Setting{}
	result := s.DB(ctx).Model(&model.Setting{}).Where("key = ?", key).Last(setting)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return setting, nil
}

func (s *Setting) Update(ctx context.Context, key string, update interface{}) (*model.Setting, error) {
	setting := &model.Setting{}
	result := s.DB(ctx).Model(&model.Setting{}).Where("key = ?", key).Update("data", update)
	if result.Error != nil {
		return nil, result.Error
	}
	return setting, nil
}
