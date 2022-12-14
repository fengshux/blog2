package model

import (
	"encoding/json"
	"time"
)

type Setting struct {
	Key        string          `gorm:"column:key;primaryKey" json:"key"`
	Data       json.RawMessage `gorm:"column:data" json:"data"`
	CreateTime time.Time       `gorm:"column:create_time;->" json:"create_time"`
	UpdateTime time.Time       `gorm:"column:update_time;->" json:"update_time"`
}

var (
	SETTING_KEY_ABOUT string = "about"
)

type SettingAbout struct {
	Content string `gorm:"column:Content" json:"content"`
}
