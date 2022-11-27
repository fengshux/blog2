package model

import "time"

type User struct {
	ID         int64     `gorm:"column:id" json:"id"`
	UserName   string    `gorm:"column:username" json:"username"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	Email      string    `gorm:"column:email" json:"email"`
	Role       string    `gorm:"column:role" json:"role"`
	CreateTime time.Time `gorm:"column:create_time;->" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;->" json:"update_time"`
}
