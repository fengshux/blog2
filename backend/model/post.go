package model

import "time"

// 文章是依懒其它数据，因此文章可以硬删除
type Post struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	Title      string    `gorm:"column:title" json:"title"`
	Body       string    `gorm:"column:body" json:"body"`
	Status     string    `gorm:"column:status" json:"status"` // draft, private, published
	TagIds     []int64   `gorm:"column:tag_ids" json:"tag_ids"`
	UserId     int64     `gorm:"column:user_id" json:"user_id"`
	CreateTime time.Time `gorm:"column:create_time;->" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;->" json:"update_time"`
}
