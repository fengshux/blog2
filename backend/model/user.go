package model

import "time"

type User struct {
	ID         int64      `gorm:"column:id"`
	UserName   string     `gorm:"column:username"`
	Email      string     `gorm:"column:email"`
	Role       string     `gorm:"column:role"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
}
