package models

import "time"

type DatabaseModels struct {
	ID          int32  `gorm:"primaryKey"`
	TaskID      int32  `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"'`
	isCompleted bool   `gorm:"default:false"`
	Done        bool   `gorm:"default:false"`
	Created_at  time.Time
}
