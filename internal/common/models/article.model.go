package models

import "time"

// uint64
// Số nguyên không dấu

// int64
// Số nguyên có dâu

type Article struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   *string   `gorm:"type:varchar(255)"`
	ImageUrl  *string   `gorm:"type:varchar(255);column:image_url"`
	LikeCount int64     `gorm:"default:0;column:like_count"`
	Views     int64     `gorm:"default:0"`
	UserID    uint64    `gorm:"not null;column:user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at"`
}
