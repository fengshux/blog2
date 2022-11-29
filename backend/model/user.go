package model

import "time"

// 用户注消也可以硬删除，删除之后，用户主页不可访问， 文章中显示用户已注消
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
