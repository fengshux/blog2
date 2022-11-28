package model

import "time"

type User struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	UserName   string    `gorm:"column:username" json:"username"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	Email      string    `gorm:"column:email" json:"email"`
	Gender     string    `gorm:"column:gender" json:"gender"`
	Role       string    `gorm:"column:role" json:"role"`
	Password   string    `gorm:"column:password" json:"password"`
	CreateTime time.Time `gorm:"column:create_time;->" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;->" json:"update_time"`
}
