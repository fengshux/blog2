package model

/**
 * 文章分类
 * 每个用户的分类都是自己的，因此分类上有userId做区分
 * 删除分类时用物理删除，同时把使此分类的文章的分类置为null
 * 关于性能，因为一个分类只能被一个同户使用，分类下的文章的数
 * 量正常会在千以下的级別，所以删除时，更新没有性能问题。
 * */

import "time"

type Category struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	UserId     int64     `gorm:"column:user_id" json:"user_id"`
	CreateTime time.Time `gorm:"column:create_time;->" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;->" json:"update_time"`
}
