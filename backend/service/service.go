package service

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type BaseService struct {
	db *gorm.DB
}

func NewBaseService(db *gorm.DB) BaseService {
	return BaseService{
		db: db,
	}
}

func (b BaseService) DB(ctx context.Context) *gorm.DB {
	return b.db.WithContext(ctx).Debug()
}

// ProviderSet is controller providers.
var ProviderSet = wire.NewSet(
	NewBaseService,
	NewUser,
	NewPost,
)
